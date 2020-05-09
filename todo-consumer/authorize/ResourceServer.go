package authorize

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
)

//设置路径权限  权限为$$LoginWhite$$登录用户放行
var PermissionPathSet = []PermissionPath{
	{
		Path:       "/api/",
		Permission: "test",
	},
	{
		Path: "/api/v1/todo/add",
		Permission: "$$LoginWhite$$",
	},
	{
		Path: "/api/v1/todo/update/**",
		Permission: "test",
	},
	{
		Path: "/api/v1/todo/delete",
		Permission: "dashboard",
	},
}
//白名单路径 不需要权限
var WhiteList = []string{
	"/api/v1/todo/all",
}

type PermissionPath struct {
	Path string 		//路径
	Permission string   //权限
}
//ResourceServerAuthorize 资源服务器鉴权 配合Spring Oauth2
func ResourceServerAuthorize() gin.HandlerFunc {
	return func(c *gin.Context){
		path := c.Request.URL.Path
		if checkWhiteList(WhiteList,path) {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		}else {



		authorization := c.GetHeader("Authorization")
		if len(authorization)<7 {
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"code":40001,"message":"访问未授权","data":nil})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
		token := authorization[7:len(authorization)]
		if checkToken(token,path) {
			// 验证通过，会继续访问下一个中间件
			//log.Info().Msg("验证通过 token: "+token)
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"code":40001,"message":"访问未授权","data":nil})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
		}
  }
}

func checkWhiteList(whiteList []string,path string)(status bool){
	for _, s := range whiteList {
		if m, _ := filepath.Match(s, path); m {
			return true
		}
	}
	return false
}

func checkToken(token string,path string)(status bool)  {
	checkTokenUrl := viper.GetString("check.token.url")
	response, err := http.Get(checkTokenUrl+"?token="+token)
	if err != nil {
		log.Error().Err(err).Msg("访问鉴权接口失败，请检查鉴权接口是否正常")
		return false
	}
	defer response.Body.Close()
	if response.StatusCode!=200 {
		return false
	}
	body, _ := ioutil.ReadAll(response.Body)
	strbody := string(body)
	var dat map[string]interface{}
	err = json.Unmarshal([]byte(strbody), &dat);
	if err!=nil {
		log.Error().Err(err).Caller()
		return false
	}
	of := reflect.ValueOf(dat["authorities"])
	//paths := []PermissionPath{
	//	{
	//		Path:       "/api/*",
	//		Permission: "test",
	//	},
	//}
	
	if !checkPath(of,PermissionPathSet,path) {
		return false
	}

	return true
}

func checkPath(of reflect.Value,pmsPath []PermissionPath,path string )(status bool)  {
	for _, permissionPath := range pmsPath {
		count := of.Len()
		for i := 0; i < count; i++ {
			pms := of.Index(i).Interface().(string)
			if permissionPath.Permission == pms||permissionPath.Permission == "$$LoginWhite$$" {
				if m, _ := filepath.Match(permissionPath.Path, path); m {
					return true
				}
			}
		}
	}
	return false
}
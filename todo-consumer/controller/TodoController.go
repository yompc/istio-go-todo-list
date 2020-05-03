package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yom535.coding.net/todo-consumer/service"
)

func GetAllTodo(c *gin.Context) {
	allTodo, err := service.TodoSelectService()

	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":40000,
			"message":"查询所有失败",
			"data":nil,
		})
	}else {

		type data struct {
			ID 			int64		`json:"id"`
			Title 		string 		`json:"title"`
			Status 		bool     	`json:"status"`
		}

		datas := make([]data, 0, 0)
		for _, todo := range allTodo.Todo {
			d := data{
				ID:     todo.Id,
				Title:  todo.Title,
				Status: todo.Status,
			}
			datas = append(datas, d)
		}
		c.JSON(http.StatusOK,gin.H{
			"code":20000,
			"message":"查询所有Todo成功",
			"data":datas,
		})
	}
}

func CreateTodo(c *gin.Context)  {
	type title struct {
		Title string `json:"title"`
	}
	var getTitle = title{}
	err := c.BindJSON(&getTitle)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest,gin.H{
			"code":40000,
			"message":"新建Task失败",
			"data":nil,
		})
		return
	}
	i := len(getTitle.Title)
	if i<10 {
		c.JSON(http.StatusBadRequest,gin.H{
			"code":40000,
			"message":"新建任务失败,请检查输入长度",
			"data":nil,
		})
		return
	}
	//err = models.CreateTodo(getTitle.Title)

	request := service.TodoWriteRequest{
		Todo: &service.Todo{
			Title: getTitle.Title,
		},
	}
	err = service.CreateTodo(&request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest,gin.H{
			"code":40000,
			"message":"新建Task失败",
			"data":nil,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":20000,
		"message":"新建Task成功",
		"data":nil,
	})
}


func UpdateTodoStateById(c *gin.Context)  {
	form := c.Query("id")
	result, _ := strconv.Atoi(form)
	//err := models.UpdateTodoStateById(result)
	request := service.TodoWriteRequest{
		Todo: &service.Todo{
			Id: int64(result),
		},
	}
	err := service.UpdateTodoStateById(&request)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":40000,
			"message":"修改Todo状态失败",
			"data":nil,
		})
	} else
	{
		c.JSON(http.StatusOK,gin.H{
			"code":20000,
			"message":"修改Todo状态成功",
			"data":nil,
		})
	}
}

func UpdateTodoTitleById(c *gin.Context)  {
	type title struct {
		ID int    		`json:"id"`
		Title string 	`json:"title"`
	}
	var result = title{}
	c.BindJSON(&result)
	i := len(result.Title)
	if i<10 {
		c.JSON(http.StatusBadRequest,gin.H{
			"code":40000,
			"message":"修改任务失败,请检查输入长度",
			"data":nil,
		})
		return
	}
	//err := models.UpdateTodoTitleById(result.ID, result.Title)
	request := service.TodoWriteRequest{
		Todo: &service.Todo{
			Id: int64(result.ID),
			Title: result.Title,
		},
	}
	err := service.UpdateTodoTitleById(&request)
	if err!=nil {

		c.JSON(http.StatusOK,gin.H{
			"code":40000,
			"message":"修改Todo内容失败",
			"data":nil,
		})
	} else
	{
		c.JSON(http.StatusOK,gin.H{
			"code":20000,
			"message":"修改Todo内容成功",
			"data":nil,
		})
	}
}

func DeleteTodoById(c *gin.Context)  {
	formId := c.Query("id")
	intId, _ := strconv.Atoi(formId)
	//err := models.DeleteTodoById(intId)
	request := service.TodoWriteRequest{
		Todo: &service.Todo{
			Id: int64(intId),
		},
	}
	err := service.DeleteTodoById(&request)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"code":40000,
			"message":"删除todo :"+formId+" 失败",
			"data":nil,
		})
	} else
	{
		c.JSON(http.StatusOK,gin.H{
			"code":20000,
			"message":"删除todo :"+formId+" 成功",
			"data":nil,
		})
	}
}
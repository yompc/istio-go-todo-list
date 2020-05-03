package main

import (
	"github.com/spf13/viper"
	"yom535.coding.net/todo-consumer/config"
	"yom535.coding.net/todo-consumer/router"
)

func main() {

	config.InitConfig()

	//初始化路由

	todoRouter := router.SetupTodoRouter()
	viper.BindEnv("web.port","WEB_PORT")
	webPort := viper.GetString("web.port")
	todoRouter.Run("0.0.0.0:"+webPort)
}

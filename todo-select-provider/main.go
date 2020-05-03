package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"yom535.coding.net/todo-select-provider/config"
	"yom535.coding.net/todo-select-provider/dao"
	"yom535.coding.net/todo-select-provider/models"
	"yom535.coding.net/todo-select-provider/service"
)

func main() {
	config.InitConfig()
	//MySQL数据库初始化
	err := dao.InitMySql()
	if err != nil {
		log.Error().Err(err).Caller()
	}
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	//END MySQL数据库初始化
	//开启GRPC服务端
	server := grpc.NewServer()
	service.RegisterTodoServiceServer(server,new(service.TodoService))
	rpcPort := viper.GetString("rpc.port")
	listen, _ := net.Listen("tcp", "0.0.0.0:"+rpcPort)
	log.Info().Msg("GRPC Provider start 0.0.0.0:"+rpcPort)
	err = server.Serve(listen)
	if err != nil {
		log.Error().Err(err).Caller()
	}
	//开启GRPC服务端
}

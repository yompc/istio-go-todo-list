package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"yom535.coding.net/todo-write-provider/config"
	"yom535.coding.net/todo-write-provider/dao"
	"yom535.coding.net/todo-write-provider/models"
	"yom535.coding.net/todo-write-provider/service"
)

func main() {
	//加载配置
	config.InitConfig()
	//END 加载配置

	//MySQL数据库初始化
	err := dao.InitMySQL()
	if err != nil {
		log.Error().Err(err).Caller()
	}
	defer dao.Close()
	dao.DB.AutoMigrate(&models.Todo{})
	//END MySQL数据库初始化
	//开启GRPC服务端
	server := grpc.NewServer()
	service.RegisterTodoServiceServer(server,new(service.TodoService))
	_ = viper.BindEnv("rpc.port","RPC_PORT")
	var rpcPort = viper.GetString("rpc.port")
	listen, _ := net.Listen("tcp", "0.0.0.0:"+rpcPort)
	log.Info().Msg("GRPC Provider Start 0.0.0.0:"+rpcPort)
	err = server.Serve(listen)
	log.Error().Err(err).Caller()
	//开启GRPC服务端

}

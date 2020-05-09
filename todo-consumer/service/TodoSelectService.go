package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	TodoSelectServicePort string
	TodoSelectServiceUri string
	TodoWriteServicePort string
	TodoWriteServiceUri string
)

func getConfg()  {
	TodoSelectServicePort = viper.GetString("todo.select.port")
	TodoSelectServiceUri = viper.GetString("todo.select.uri")
	TodoWriteServicePort = viper.GetString("todo.write.port")
	TodoWriteServiceUri = viper.GetString("todo.write.uri")
}

func TodoSelectService() (response *TodoResponse,err error) {
	getConfg()
	//log.Info().Caller().Msg("conn :"+TodoSelectServiceUri+":"+TodoSelectServicePort)
	TodoSelectConn, err := grpc.Dial(TodoSelectServiceUri+":"+TodoSelectServicePort,grpc.WithInsecure())
	if err!=nil {
		log.Error().Err(err)
	}
	todoClient := NewTodoServiceClient(TodoSelectConn)
	defer TodoSelectConn.Close()
	Todos, err := todoClient.TodoDTO(context.Background(),&TodoRequest{})
	return Todos,err
}

func CreateTodo(in *TodoWriteRequest)(err error)  {
	getConfg()
	TodoSelectConn, err := grpc.Dial(TodoWriteServiceUri+":"+TodoWriteServicePort,grpc.WithInsecure())
	if err!=nil {
		log.Error().Err(err)
	}
	defer TodoSelectConn.Close()
	writeServiceClient := NewTodoServiceClient(TodoSelectConn)
	_, err = writeServiceClient.CreateTodo(context.Background(), in)
	return err
}

func UpdateTodoStateById(in *TodoWriteRequest)(err error)  {
	getConfg()
	TodoSelectConn, err := grpc.Dial(TodoWriteServiceUri+":"+TodoWriteServicePort,grpc.WithInsecure())
	if err!=nil {
		log.Error().Err(err)
	}
	defer TodoSelectConn.Close()
	writeServiceClient := NewTodoServiceClient(TodoSelectConn)
	_, err = writeServiceClient.UpdateTodoStateById(context.Background(), in)
	return err
}

func UpdateTodoTitleById(in *TodoWriteRequest)(err error)  {
	getConfg()
	TodoSelectConn, err := grpc.Dial(TodoWriteServiceUri+":"+TodoWriteServicePort,grpc.WithInsecure())
	if err!=nil {
		log.Error().Err(err)
	}
	defer TodoSelectConn.Close()
	writeServiceClient := NewTodoServiceClient(TodoSelectConn)
	_, err = writeServiceClient.UpdateTodoTitleById(context.Background(), in)
	return err
}

func DeleteTodoById(in *TodoWriteRequest)(err error)   {
	getConfg()
	TodoSelectConn, err := grpc.Dial(TodoWriteServiceUri+":"+TodoWriteServicePort,grpc.WithInsecure())
	if err!=nil {
		log.Error().Err(err)
	}
	defer TodoSelectConn.Close()
	writeServiceClient := NewTodoServiceClient(TodoSelectConn)
	_, err = writeServiceClient.DeleteTodoById(context.Background(), in)
	return err
}
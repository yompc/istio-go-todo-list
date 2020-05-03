package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"yom535.coding.net/todo-write-provider/models"
)

type TodoService struct {

}

func (this *TodoService) TodoDTO(ctx context.Context, request *TodoRequest) (*TodoResponse, error) {
	panic("implement me")
}

func (this *TodoService) CreateTodo(ctx context.Context, in *TodoWriteRequest) (*TodoWriteResponse, error)  {
	todo := in.Todo
	err := models.CreateTodo(todo.Title)
	if err != nil {
		log.Error().Err(err).Caller()
	}
	return &TodoWriteResponse{Status: 1},err
}

func (this *TodoService)UpdateTodoStateById(ctx context.Context,in *TodoWriteRequest) (*TodoWriteResponse, error)  {
	todo := in.Todo
	err := models.UpdateTodoStateById(todo.Id)
	if err != nil {
		log.Error().Err(err).Caller()
	}
	return &TodoWriteResponse{Status: 1},err
}

func (this *TodoService)UpdateTodoTitleById(cex context.Context,in *TodoWriteRequest) (*TodoWriteResponse, error)  {
	todo := in.Todo
	err := models.UpdateTodoTitleById(int(todo.Id), todo.Title)
	if err != nil {
		log.Error().Err(err).Caller()
	}
	return &TodoWriteResponse{Status: 1},err
}

func (this *TodoService)DeleteTodoById(ctx context.Context,in *TodoWriteRequest) (*TodoWriteResponse, error) {
	todo := in.Todo
	err := models.DeleteTodoById(todo.Id)
	if err != nil {
		log.Error().Err(err).Caller()
	}
	return &TodoWriteResponse{Status: 1},err
}



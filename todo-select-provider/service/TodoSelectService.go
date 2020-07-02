package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"yom535.coding.net/todo-select-provider/models"
)

type TodoService struct {

}

func (this *TodoService) CreateTodo(ctx context.Context, request *TodoWriteRequest) (*TodoWriteResponse, error) {
	panic("implement me")
}

func (this *TodoService) UpdateTodoStateById(ctx context.Context, request *TodoWriteRequest) (*TodoWriteResponse, error) {
	panic("implement me")
}

func (this *TodoService) UpdateTodoTitleById(ctx context.Context, request *TodoWriteRequest) (*TodoWriteResponse, error) {
	panic("implement me")
}

func (this *TodoService) DeleteTodoById(ctx context.Context, request *TodoWriteRequest) (*TodoWriteResponse, error) {
	panic("implement me")
}

func (this *TodoService) TodoDTO(ctx context.Context,request *TodoRequest) (*TodoResponse, error) {
	listTodo, err := models.GetAllListTodo()
	if err != nil {
		log.Error().Err(err).Caller()
	}
	var response = []*Todo{}
	for _, todo := range listTodo {
		t := &Todo{
			Id:     todo.ID,
			Title:  todo.Title,
			Status: todo.Status,
		}
		response = append(response, t)
	}
	return &TodoResponse{
		Todo: response,
	},nil
}

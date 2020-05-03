package models

import (
	"github.com/rs/zerolog/log"
	"yom535.coding.net/todo-select-provider/dao"
)

type Todo struct {
	ID 			int64		`json:"id"`
	Title 		string 		`json:"title"`
	Status 		bool     	`json:"status"`
}

func (Todo) TableName() string {
	return "todo"
}
/**
	获取所有Todo
	@Return todoList 所有Todo
 */
func GetAllListTodo()(todoList []*Todo,err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		log.Error().Err(err).Caller()
	}
	return
}
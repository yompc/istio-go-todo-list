package models

import "yom535.coding.net/todo-write-provider/dao"

type Todo struct {
	ID 			int64			`json:"id"`
	Title 		string 		`json:"title"`
	Status 		bool     	`json:"status"`
}

func (Todo) TableName() string {
	return "todo"
}

/**
创建Todo
@Param title  标题
@Return err
*/
func CreateTodo(title string) (err error) {
	var todo = Todo{
		Title: title,
		Status: false,
	}
	err = dao.DB.Save(&todo).Error
	return
}
/**
更新状态
@Param id ID
@Return err
*/
func UpdateTodoStateById(id int64) (err error) {
	var todo =Todo{
		ID:    id,
		Status: true,
	}
	err = dao.DB.Model(&Todo{}).Update(&todo).Error
	return
}
/**
更新标题
@Param id ID
@Param title 标题
*/
func UpdateTodoTitleById(id int , title string)(err error)  {
	err = dao.DB.Exec("UPDATE todo SET `title`=? , `status` = 0 WHERE id = ?", title, id).Error
	return
}

/**
根据ID删除Todo
@Param id ID
@Return err
*/
func DeleteTodoById(id int64) (err error) {
	if id==0 {
		return
	}
	var todo = Todo{
		ID:    id,
	}
	err = dao.DB.Delete(&todo).Error
	return
}

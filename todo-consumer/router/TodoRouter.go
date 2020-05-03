package router

import (
	"github.com/gin-gonic/gin"
	"yom535.coding.net/todo-consumer/controller"
)

func SetupTodoRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	group := r.Group("/api/v1/todo")
	{
		group.GET("/all",controller.GetAllTodo)
		group.POST("/add",controller.CreateTodo)
		group.POST("/update/status",controller.UpdateTodoStateById)
		group.POST("/update/title",controller.UpdateTodoTitleById)
		group.POST("/delete",controller.DeleteTodoById)
	}
	return r
}

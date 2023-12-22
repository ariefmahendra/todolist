package router

import (
	"github.com/gin-gonic/gin"
	"todolist/controller"
)

func NewRouter(todolistController controller.TodolistController) *gin.Engine {
	r := gin.Default()

	r.GET("/api/todolist", todolistController.FindAll)
	r.POST("/api/todolist", todolistController.Create)
	r.PUT("/api/todolist/:id", todolistController.Update)
	r.DELETE("/api/todolist/:id", todolistController.Delete)
	r.GET("/api/todolist/:id", todolistController.FindById)

	return r
}

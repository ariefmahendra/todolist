package router

import (
	"github.com/gin-gonic/gin"
	"todolist/controller"
	"todolist/exception"
)

func NewRouter(todolistController controller.TodolistController, userController controller.UserController) *gin.Engine {
	r := gin.Default()

	r.Use(gin.CustomRecovery(exception.ErrorHandler))

	r.POST("/auth/register", userController.Register)
	r.POST("/auth/login", userController.Login)

	r.GET("/api/todolist", todolistController.FindAll)
	r.POST("/api/todolist", todolistController.Create)
	r.PUT("/api/todolist/:id", todolistController.Update)
	r.DELETE("/api/todolist/:id", todolistController.Delete)
	r.GET("/api/todolist/:id", todolistController.FindById)

	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	"todolist/controller"
	"todolist/exception"
	"todolist/middleware"
)

func NewRouter(todolistController controller.TodolistController, userController controller.UserController, scopesController controller.ScopeController, assignController controller.AssignmentController) *gin.Engine {
	r := gin.Default()

	r.Use(gin.CustomRecovery(exception.ErrorHandler))
	r.Use(middleware.AuthMiddleware)

	r.POST("/auth/register", userController.Register)
	r.POST("/auth/login", userController.Login)

	r.POST("/api/scopes", scopesController.Create)
	r.DELETE("/api/scopes/:name", scopesController.Delete)

	r.GET("/api/user/:user_id/scopes/:scope_id/assign", assignController.AssignUserScope)

	r.GET("/api/todolist", todolistController.FindAll)
	r.POST("/api/todolist", todolistController.Create)
	r.PUT("/api/todolist/:id", todolistController.Update)
	r.DELETE("/api/todolist/:id", todolistController.Delete)
	r.GET("/api/todolist/:id", todolistController.FindById)
	r.PUT("/api/todolist/:id/check", todolistController.CheckTodolist)

	return r
}

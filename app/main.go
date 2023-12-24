package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
	"todolist/config"
	"todolist/controller"
	"todolist/helper"
	"todolist/repository"
	"todolist/router"
	"todolist/service"
)

func main() {

	db := config.InitializedDatabase()
	validate := validator.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	todolistRepository := repository.NewTodolistRepository(db)
	todolistService := service.NewTodolistService(todolistRepository, db, validate)
	todolistController := controller.NewTodolistController(todolistService)

	route := router.NewRouter(todolistController, userController)

	server := http.Server{
		Addr:           ":8080",
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

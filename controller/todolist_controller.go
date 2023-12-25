package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/exception"
	"todolist/helper"
	"todolist/model/dto"
	"todolist/model/middleware"
	"todolist/service"
)

type TodolistController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	CheckTodolist(ctx *gin.Context)
}

type TodolistControllerImpl struct {
	todolistService service.TodolistService
}

func NewTodolistController(todolistService service.TodolistService) TodolistController {
	return &TodolistControllerImpl{todolistService: todolistService}
}

func (controller *TodolistControllerImpl) Create(ctx *gin.Context) {
	createdTodolist := dto.CreateTodolistRequest{}

	err := ctx.BindJSON(&createdTodolist)
	helper.PanicIfError(err)

	value, exists := ctx.Get("currentUser")
	if !exists {
		panic(exception.NewUnauthorizedError("User not found"))
	}

	currentUser := value.(middleware.AuthClaimJWT)
	createdTodolist.UserId = currentUser.UserId

	for _, scope := range currentUser.Scopes {
		if scope == "todos:create" {
			todolistResponse := controller.todolistService.Create(ctx, createdTodolist)

			success := dto.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   todolistResponse,
			}

			ctx.JSON(http.StatusOK, success)

			return
		}
	}

	panic(exception.NewForbiddenError("You don't have permission to create todolist"))
}

func (controller *TodolistControllerImpl) Update(ctx *gin.Context) {
	updatedTodolist := dto.UpdateTodolistRequest{}

	todolistId := ctx.Param("id")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	err = ctx.BindJSON(&updatedTodolist)
	helper.PanicIfError(err)

	updatedTodolist.Id = id

	value, exists := ctx.Get("currentUser")
	if !exists {
		panic(exception.NewUnauthorizedError("User not found"))
	}

	currentUser := value.(middleware.AuthClaimJWT)
	updatedTodolist.UserId = currentUser.UserId

	for _, scope := range currentUser.Scopes {
		if scope == "todos:update" {
			todolistResponse := controller.todolistService.Update(ctx, updatedTodolist)

			response := dto.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   todolistResponse,
			}

			ctx.JSON(http.StatusOK, response)
			return
		}
	}
	panic(exception.NewForbiddenError("You don't have permission to update todolist"))
}

func (controller *TodolistControllerImpl) Delete(ctx *gin.Context) {
	deletedTodolist := dto.TodolistResponse{}

	todolistId := ctx.Param("id")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	var userId int
	value, exists := ctx.Get("currentUser")
	if !exists {
		panic(exception.NewUnauthorizedError("User not found"))
	}

	currentUser := value.(middleware.AuthClaimJWT)
	userId = currentUser.UserId

	for _, scope := range currentUser.Scopes {
		if scope == "todos:delete" {
			controller.todolistService.Delete(ctx, id, userId)

			response := dto.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   deletedTodolist,
			}

			ctx.JSON(http.StatusOK, response)
			return
		}
	}
	panic(exception.NewForbiddenError("You don't have permission to delete todolist"))
}

func (controller *TodolistControllerImpl) FindAll(ctx *gin.Context) {
	var userId int
	value, exists := ctx.Get("currentUser")
	if !exists {
		panic(exception.NewUnauthorizedError("User not found"))
	}

	currentUser := value.(middleware.AuthClaimJWT)
	userId = currentUser.UserId

	for _, scope := range currentUser.Scopes {
		if scope == "todos:read" {
			todolistResponse := controller.todolistService.GetAll(ctx, userId)

			response := dto.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   todolistResponse,
			}

			ctx.JSON(http.StatusOK, response)
			return
		}
	}
	panic(exception.NewForbiddenError("You don't have permission to read todolist"))
}

func (controller *TodolistControllerImpl) FindById(ctx *gin.Context) {
	todolistId := ctx.Param("id")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	var userId int
	value, exists := ctx.Get("currentUser")
	if !exists {
		panic(exception.NewUnauthorizedError("User not found"))
	}

	currentUser := value.(middleware.AuthClaimJWT)
	userId = currentUser.UserId

	for _, scope := range currentUser.Scopes {
		if scope == "todos:read" {
			todolistResponse := controller.todolistService.GetById(ctx, id, userId)

			response := dto.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   todolistResponse,
			}

			ctx.JSON(http.StatusOK, response)
			return
		}
	}
	panic(exception.NewForbiddenError("You don't have permission to read todolist"))
}

func (controller *TodolistControllerImpl) CheckTodolist(ctx *gin.Context) {
	todolistId := ctx.Param("id")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	var checkRequest dto.CheckTodolistRequest
	err = ctx.BindJSON(&checkRequest)
	if err != nil {
		panic("Error when binding JSON")
	}

	var userId int
	value, exists := ctx.Get("currentUser")
	if !exists {
		panic(exception.NewUnauthorizedError("User not found"))
	}

	currentUser := value.(middleware.AuthClaimJWT)
	userId = currentUser.UserId

	for _, scope := range currentUser.Scopes {
		if scope == "todos:update" {
			todolistResponse := controller.todolistService.CheckTodolist(ctx, checkRequest, id, userId)

			response := dto.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   todolistResponse,
			}

			ctx.JSON(http.StatusOK, response)
			return
		}
	}
	panic(exception.NewForbiddenError("You don't have permission to check todolist"))
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/helper"
	"todolist/model/dto"
	"todolist/service"
)

type TodolistController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
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

	todolistResponse := controller.todolistService.Create(ctx, createdTodolist)

	success := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todolistResponse,
	}

	ctx.JSON(http.StatusOK, success)
}

func (controller *TodolistControllerImpl) Update(ctx *gin.Context) {
	updatedTodolist := dto.UpdateTodolistRequest{}

	todolistId := ctx.Param("id")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	err = ctx.BindJSON(&updatedTodolist)
	helper.PanicIfError(err)

	updatedTodolist.Id = id

	todolistResponse := controller.todolistService.Update(ctx, updatedTodolist)

	response := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todolistResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *TodolistControllerImpl) Delete(ctx *gin.Context) {
	deletedTodolist := dto.TodolistResponse{}

	todolistId := ctx.Param("id")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	controller.todolistService.Delete(ctx, id)

	response := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   deletedTodolist,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *TodolistControllerImpl) FindAll(ctx *gin.Context) {
	todolistResponse := controller.todolistService.GetAll(ctx)

	response := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todolistResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *TodolistControllerImpl) FindById(ctx *gin.Context) {
	todolistId := ctx.Param("id")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	todolistResponse := controller.todolistService.GetById(ctx, id)

	response := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todolistResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

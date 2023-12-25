package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/helper"
	"todolist/model/dto"
	"todolist/service"
)

type ScopeController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ScopeControllerImpl struct {
	scopesService service.ScopeService
}

func NewScopeController(scopesService service.ScopeService) ScopeController {
	return &ScopeControllerImpl{scopesService: scopesService}
}

func (controller *ScopeControllerImpl) Create(ctx *gin.Context) {
	var scopesCreated dto.CreateScopeRequest
	err := ctx.BindJSON(&scopesCreated)
	helper.PanicIfError(err)

	scopesResponse := controller.scopesService.Create(ctx, scopesCreated)

	response := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   scopesResponse,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *ScopeControllerImpl) Delete(ctx *gin.Context) {
	name := ctx.Param("name")

	controller.scopesService.Delete(ctx, name)

	response := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, response)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/helper"
	"todolist/model/dto"
	"todolist/service"
)

type AssignmentController interface {
	AssignUserScope(ctx *gin.Context)
}

type AssignmentControllerImpl struct {
	assignmentService service.AssignmentService
}

func NewAssignmentController(assignmentService service.AssignmentService) AssignmentController {
	return &AssignmentControllerImpl{assignmentService: assignmentService}
}

func (controller *AssignmentControllerImpl) AssignUserScope(ctx *gin.Context) {

	var request dto.AssignmentUserScopeRequest

	paramUserId := ctx.Param("user_id")
	userId, err := strconv.Atoi(paramUserId)
	helper.PanicIfError(err)

	paramScopeId := ctx.Param("scope_id")
	scopeId, err := strconv.Atoi(paramScopeId)
	helper.PanicIfError(err)

	request.UserId = userId
	request.ScopeId = scopeId

	assignResponse := controller.assignmentService.Assign(ctx, request)

	response := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   assignResponse,
	}

	ctx.JSON(200, response)
}

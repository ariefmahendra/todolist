package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/helper"
	"todolist/model/dto"
	"todolist/service"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{userService: userService}
}

func (controller *UserControllerImpl) Register(ctx *gin.Context) {
	var userRegisterRequest dto.CreateUserRequest

	err := ctx.BindJSON(&userRegisterRequest)
	helper.PanicIfError(err)

	userResponse := controller.userService.Create(ctx, userRegisterRequest)

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Login(ctx *gin.Context) {
	var userLoginRequest dto.LoginRequest

	err := ctx.BindJSON(&userLoginRequest)
	helper.PanicIfError(err)

	token, scopes := controller.userService.Login(ctx, userLoginRequest)

	bearerToken := fmt.Sprintf("Bearer %s", token)

	loginResponse := dto.LoginUserResponse{
		Email:  userLoginRequest.Email,
		Token:  bearerToken,
		Scopes: scopes,
	}

	webResponse := dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   loginResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

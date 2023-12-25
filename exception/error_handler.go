package exception

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"todolist/model/dto"
)

func ErrorHandler(ctx *gin.Context, err interface{}) {
	if validationErrors(ctx, err) {
		return
	}

	if notFoundError(ctx, err) {
		return
	}

	if unauthorizedError(ctx, err) {
		return
	}

	if forbiddenError(ctx, err) {
		return
	}

	internal(ctx, err)

}

func forbiddenError(ctx *gin.Context, err interface{}) bool {
	exception, ok := err.(ForbiddenError)
	if ok {
		response := dto.WebResponse{
			Code:   http.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   exception.Error,
		}

		ctx.JSON(http.StatusForbidden, response)
		return true
	}
	return false
}

func unauthorizedError(ctx *gin.Context, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		response := dto.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}

		ctx.JSON(http.StatusUnauthorized, response)
		return true
	}
	return false
}

func validationErrors(ctx *gin.Context, err interface{}) bool {
	var exception validator.ValidationErrors
	if e, ok := err.(error); ok {
		if errors.As(e, &exception) {
			response := dto.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   exception.Error(),
			}

			ctx.JSON(http.StatusBadRequest, response)
			return true
		}
	}
	return false
}

func notFoundError(ctx *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		response := dto.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		ctx.JSON(http.StatusNotFound, response)
		return true
	}
	return false
}

func internal(ctx *gin.Context, err interface{}) {
	response := dto.WebResponse{
		Code:   500,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	ctx.JSON(http.StatusInternalServerError, response)
}

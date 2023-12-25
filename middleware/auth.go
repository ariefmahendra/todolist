package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"todolist/exception"
	"todolist/model/middleware"
)

func AuthMiddleware(ctx *gin.Context) {
	if ctx.FullPath() != "/auth/login" && ctx.FullPath() != "/auth/register" {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			panic(exception.NewUnauthorizedError("empty token"))
		}

		authArr := strings.Split(authHeader, " ")
		if len(authArr) != 2 {
			panic(exception.NewUnauthorizedError("invalid token"))
		}

		if authArr[0] != "Bearer" {
			panic(exception.NewUnauthorizedError("invalid bearer token"))
		}

		tokenStr := authArr[1]

		var authClaim middleware.AuthClaimJWT

		token, err := jwt.ParseWithClaims(tokenStr, &authClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte("123"), nil
		})

		if err != nil {
			panic(err.Error())
		}

		if !token.Valid {
			panic(exception.NewUnauthorizedError("invalid token"))
		}

		ctx.Set("currentUser", authClaim)

		ctx.Next()
	}
}

package middleware

import "github.com/golang-jwt/jwt/v5"

type AuthClaimJWT struct {
	UserId int      `json:"user_id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Scopes []string `json:"scopes"`
	jwt.RegisteredClaims
}

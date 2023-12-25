package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"todolist/helper"
	"todolist/model/domain"
	"todolist/model/dto"
	"todolist/model/middleware"
	"todolist/repository"
)

type UserService interface {
	Create(ctx context.Context, user dto.CreateUserRequest) dto.UserResponse
	Login(ctx context.Context, user dto.LoginRequest) (string, []string)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
	db             *sql.DB
	validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{userRepository: userRepository, db: db, validate: validate}
}

func (service *UserServiceImpl) Create(ctx context.Context, user dto.CreateUserRequest) dto.UserResponse {
	err := service.validate.Struct(user)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	userDomain := domain.UserDomain{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(password),
	}

	userResponse := service.userRepository.Insert(ctx, tx, userDomain)

	return helper.UserToResponse(userResponse)
}

func (service *UserServiceImpl) Login(ctx context.Context, user dto.LoginRequest) (string, []string) {
	err := service.validate.Struct(user)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userDomain, err := service.userRepository.FindByEmail(ctx, tx, user.Email)
	helper.PanicIfError(err)

	scopes := service.userRepository.FindScopeByEmail(ctx, tx, user.Email)

	err = bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(user.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			panic("password invalid")
		}
	}

	tokenClaim := middleware.AuthClaimJWT{
		UserId: userDomain.Id,
		Name:   userDomain.Name,
		Email:  userDomain.Email,
		Scopes: scopes,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

	tokenString, err := token.SignedString([]byte("123"))

	if err != nil {
		panic("failed to generate token")
	}

	return tokenString, scopes
}

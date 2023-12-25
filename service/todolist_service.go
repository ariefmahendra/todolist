package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"todolist/exception"
	"todolist/helper"
	"todolist/model/domain"
	"todolist/model/dto"
	"todolist/repository"
)

type TodolistService interface {
	Create(ctx context.Context, request dto.CreateTodolistRequest) dto.TodolistResponse
	Update(ctx context.Context, request dto.UpdateTodolistRequest) dto.TodolistResponse
	Delete(ctx context.Context, id int, userId int)
	GetAll(ctx context.Context, userId int) []dto.TodolistResponse
	GetById(ctx context.Context, id int, userId int) dto.TodolistResponse
	CheckTodolist(ctx context.Context, request dto.CheckTodolistRequest, id int, userId int) dto.TodolistResponse
}

type TodolistServiceImpl struct {
	todolistRepository repository.TodolistRepository
	db                 *sql.DB
	validate           *validator.Validate
}

func NewTodolistService(todolistRepository repository.TodolistRepository, db *sql.DB, validate *validator.Validate) TodolistService {
	return &TodolistServiceImpl{
		todolistRepository: todolistRepository,
		db:                 db,
		validate:           validate,
	}
}

func (service *TodolistServiceImpl) Create(ctx context.Context, request dto.CreateTodolistRequest) dto.TodolistResponse {

	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist := domain.TodolistDomain{
		Title:       request.Title,
		Description: request.Description,
		IsDone:      request.IsDone,
		UserId:      request.UserId,
	}

	todolistResponse := service.todolistRepository.Insert(ctx, tx, todolist)

	return helper.TodolistToResponse(todolistResponse)
}

func (service *TodolistServiceImpl) Update(ctx context.Context, request dto.UpdateTodolistRequest) dto.TodolistResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.todolistRepository.FindById(ctx, tx, request.Id, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todolistUpdated := domain.TodolistDomain{
		Id:          todolist.Id,
		Title:       request.Title,
		Description: request.Description,
		IsDone:      request.IsDone,
		UserId:      request.UserId,
	}

	todolistUpdatedResponse := service.todolistRepository.Update(ctx, tx, todolistUpdated)

	return helper.TodolistToResponse(todolistUpdatedResponse)
}

func (service *TodolistServiceImpl) Delete(ctx context.Context, id int, userId int) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.todolistRepository.FindById(ctx, tx, id, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.todolistRepository.Delete(ctx, tx, todolist.Id, userId)
}

func (service *TodolistServiceImpl) GetAll(ctx context.Context, userId int) []dto.TodolistResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	items, err := service.todolistRepository.FindAll(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.TodolistToResponses(items)
}

func (service *TodolistServiceImpl) GetById(ctx context.Context, id int, userId int) dto.TodolistResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.todolistRepository.FindById(ctx, tx, id, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.TodolistToResponse(todolist)
}

func (service *TodolistServiceImpl) CheckTodolist(ctx context.Context, request dto.CheckTodolistRequest, id int, userId int) dto.TodolistResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.todolistRepository.CheckTodolist(ctx, tx, request, id, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.TodolistToResponse(todolist)
}

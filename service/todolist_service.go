package service

import (
	"context"
	"todolist/model/dto"
)

type TodolistService interface {
	Create(ctx context.Context, request dto.CreateTodolistRequest) dto.TodolistResponse
	Update(ctx context.Context, request dto.UpdateTodolistRequest) dto.TodolistResponse
	Delete(ctx context.Context, id int)
	GetAll(ctx context.Context) []dto.TodolistResponse
	GetById(ctx context.Context, id int) dto.TodolistResponse
}

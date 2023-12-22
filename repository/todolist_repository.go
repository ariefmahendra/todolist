package repository

import (
	"context"
	"database/sql"
	"todolist/model/domain"
)

type TodolistRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain
	Update(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TodolistDomain, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TodolistDomain, error)
}

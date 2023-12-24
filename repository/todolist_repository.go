package repository

import (
	"context"
	"database/sql"
	"errors"
	"todolist/helper"
	"todolist/model/domain"
)

type TodolistRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain
	Update(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TodolistDomain, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TodolistDomain, error)
}

type TodolistRepositoryImpl struct {
	db *sql.DB
}

func NewTodolistRepository(db *sql.DB) TodolistRepository {
	return &TodolistRepositoryImpl{db: db}
}

func (repository *TodolistRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain {
	SQL := "INSERT INTO mst_todolist (title, description) VALUES ($1, $2) RETURNING id, isDone"

	err := tx.QueryRowContext(ctx, SQL, todolistDomain.Title, todolistDomain.Description).Scan(&todolistDomain.Id, &todolistDomain.IsDone)
	helper.PanicIfError(err)

	return todolistDomain
}

func (repository *TodolistRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain {
	SQL := "update mst_todolist set title = $1, description = $2, isDone = $3 where id = $4"

	_, err := tx.ExecContext(ctx, SQL, todolistDomain.Title, todolistDomain.Description, todolistDomain.IsDone, todolistDomain.Id)
	helper.PanicIfError(err)

	return todolistDomain
}

func (repository *TodolistRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "delete from mst_todolist where id = $1"

	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (repository *TodolistRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.TodolistDomain, error) {
	SQL := "select id, title, description, isDone from mst_todolist"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var todolist []domain.TodolistDomain

	defer func() {
		err := rows.Close()
		helper.PanicIfError(err)
	}()
	for rows.Next() {
		todoItem := domain.TodolistDomain{}
		err := rows.Scan(&todoItem.Id, &todoItem.Title, &todoItem.Description, &todoItem.IsDone)
		helper.PanicIfError(err)
		todolist = append(todolist, todoItem)
	}

	if len(todolist) == 0 {
		return todolist, errors.New("todolist not found")
	}

	return todolist, nil
}

func (repository *TodolistRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.TodolistDomain, error) {
	SQL := "SELECT id, title, description, isDone FROM mst_todolist WHERE id = $1"

	todolist := domain.TodolistDomain{}

	err := tx.QueryRowContext(ctx, SQL, id).Scan(&todolist.Id, &todolist.Title, &todolist.Description, &todolist.IsDone)
	helper.PanicIfError(err)

	return todolist, nil
}

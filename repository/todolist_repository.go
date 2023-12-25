package repository

import (
	"context"
	"database/sql"
	"errors"
	"todolist/helper"
	"todolist/model/domain"
	"todolist/model/dto"
)

type TodolistRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain
	Update(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain
	Delete(ctx context.Context, tx *sql.Tx, id int, userId int)
	FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]domain.TodolistDomain, error)
	FindById(ctx context.Context, tx *sql.Tx, id int, userId int) (domain.TodolistDomain, error)
	CheckTodolist(ctx context.Context, tx *sql.Tx, request dto.CheckTodolistRequest, id int, userId int) (domain.TodolistDomain, error)
}

type TodolistRepositoryImpl struct {
	db *sql.DB
}

func NewTodolistRepository(db *sql.DB) TodolistRepository {
	return &TodolistRepositoryImpl{db: db}
}

func (repository *TodolistRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain {
	SQL := "INSERT INTO mst_todolist (title, description, user_id) VALUES ($1, $2, $3) RETURNING id, isDone"

	err := tx.QueryRowContext(ctx, SQL, todolistDomain.Title, todolistDomain.Description, todolistDomain.UserId).Scan(&todolistDomain.Id, &todolistDomain.IsDone)
	helper.PanicIfError(err)

	return todolistDomain
}

func (repository *TodolistRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todolistDomain domain.TodolistDomain) domain.TodolistDomain {
	SQL := "update mst_todolist set title = $1, description = $2, isDone = $3 where id = $4 and user_id = $5"

	_, err := tx.ExecContext(ctx, SQL, todolistDomain.Title, todolistDomain.Description, todolistDomain.IsDone, todolistDomain.Id, todolistDomain.UserId)
	helper.PanicIfError(err)

	return todolistDomain
}

func (repository *TodolistRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int, userId int) {
	SQL := "delete from mst_todolist where id = $1 and user_id = $2"

	_, err := tx.ExecContext(ctx, SQL, id, userId)
	helper.PanicIfError(err)
}

func (repository *TodolistRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]domain.TodolistDomain, error) {
	SQL := "select id, title, description, isDone from mst_todolist where user_id = $1"

	rows, err := tx.QueryContext(ctx, SQL, userId)
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

func (repository *TodolistRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int, userId int) (domain.TodolistDomain, error) {
	SQL := "SELECT id, title, description, isDone FROM mst_todolist WHERE id = $1 AND user_id = $2"

	todolist := domain.TodolistDomain{}

	err := tx.QueryRowContext(ctx, SQL, id, userId).Scan(&todolist.Id, &todolist.Title, &todolist.Description, &todolist.IsDone)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

func (repository *TodolistRepositoryImpl) CheckTodolist(ctx context.Context, tx *sql.Tx, request dto.CheckTodolistRequest, id int, userId int) (domain.TodolistDomain, error) {
	SQL := "update mst_todolist set isDone = $1 where id = $2 and user_id = $3 returning id, title, description, isDone"

	todolist := domain.TodolistDomain{}

	err := tx.QueryRowContext(ctx, SQL, request.IsDone, id, userId).Scan(&todolist.Id, &todolist.Title, &todolist.Description, &todolist.IsDone)
	if err != nil {
		return todolist, err
	}

	return todolist, nil
}

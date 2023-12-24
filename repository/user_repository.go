package repository

import (
	"context"
	"database/sql"
	"errors"
	"todolist/exception"
	"todolist/helper"
	"todolist/model/domain"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, domain domain.UserDomain) domain.UserDomain
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.UserDomain, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, domain domain.UserDomain) domain.UserDomain {
	SQL := "insert into mst_user (email, password, name) values ($1, $2, $3) returning id"

	err := tx.QueryRowContext(ctx, SQL, domain.Email, domain.Password, domain.Name).Scan(&domain.Id)
	helper.PanicIfError(err)

	return domain
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.UserDomain, error) {
	SQL := "select id, name, email from mst_user where email = $1"

	var userDomain domain.UserDomain

	err := tx.QueryRowContext(ctx, SQL, email).Scan(&userDomain.Id, &userDomain.Name, &userDomain.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(exception.NewNotFoundError(err.Error()))
		}

		panic(err)
	}

	return userDomain, nil
}

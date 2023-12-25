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
	FindScopeByEmail(ctx context.Context, tx *sql.Tx, email string) []string
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
	SQL := "select id, name, email, password from mst_user where email = $1"

	var userDomain domain.UserDomain

	err := tx.QueryRowContext(ctx, SQL, email).Scan(&userDomain.Id, &userDomain.Name, &userDomain.Email, &userDomain.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(exception.NewNotFoundError(err.Error()))
		}
		panic(err)
	}

	return userDomain, nil
}

func (repository *UserRepositoryImpl) FindScopeByEmail(ctx context.Context, tx *sql.Tx, email string) []string {
	SQL := `
			select mst_scopes.name as scope_name 
			from mst_user as u
			left join mst_user_scope on u.id = mst_user_scope.user_id
			join mst_scopes on mst_scopes.id = mst_user_scope.scope_id
			where email = $1
			`

	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)

	var scopes []string

	for rows.Next() {
		var scope string
		err := rows.Scan(&scope)
		helper.PanicIfError(err)

		scopes = append(scopes, scope)
	}

	return scopes
}

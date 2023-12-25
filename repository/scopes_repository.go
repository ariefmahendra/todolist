package repository

import (
	"context"
	"database/sql"
	"todolist/helper"
	"todolist/model/domain"
)

type ScopesRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, scopesDomain domain.ScopesDomain) domain.ScopesDomain
	Delete(ctx context.Context, tx *sql.Tx, name string)
}

type ScopesRepositoryImpl struct {
}

func NewScopesRepository() ScopesRepository {
	return &ScopesRepositoryImpl{}
}

func (scopes *ScopesRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, scopesDomain domain.ScopesDomain) domain.ScopesDomain {
	SQL := "INSERT INTO mst_scopes (name) VALUES ($1) RETURNING id"
	err := tx.QueryRowContext(ctx, SQL, scopesDomain.Name).Scan(&scopesDomain.Id)
	helper.PanicIfError(err)

	return scopesDomain
}

func (scopes *ScopesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, name string) {
	SQL := "delete from mst_scopes where name = $1"
	result, err := tx.ExecContext(ctx, SQL, name)
	helper.PanicIfError(err)

	affected, err := result.RowsAffected()
	helper.PanicIfError(err)

	if affected == 0 {
		panic("failed to delete scope")
	}
}

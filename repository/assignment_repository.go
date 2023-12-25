package repository

import (
	"context"
	"database/sql"
	"errors"
	"todolist/model/domain"
)

type AssignmentRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, assignmentDomain domain.AssignmentUserScope) domain.AssignmentUserScope
	Find(ctx context.Context, tx *sql.Tx, assignmentDomain domain.AssignmentUserScope) error
}

type AssignmentRepositoryImpl struct {
}

func NewAssignmentRepository() AssignmentRepository {
	return &AssignmentRepositoryImpl{}
}

func (repository *AssignmentRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, assignmentDomain domain.AssignmentUserScope) domain.AssignmentUserScope {
	SQL := "INSERT INTO mst_user_scope (user_id, scope_id) VALUES ($1, $2) RETURNING id"
	err := tx.QueryRowContext(ctx, SQL, assignmentDomain.UserId, assignmentDomain.ScopeId).Scan(&assignmentDomain.Id)
	if err != nil {
		panic("duplicate data")
	}

	return assignmentDomain
}

func (repository *AssignmentRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, assignmentDomain domain.AssignmentUserScope) error {

	SQL := "select id, user_id, scope_id from mst_user_scope where user_id = $1 and scope_id = $2"
	err := tx.QueryRowContext(ctx, SQL, assignmentDomain.UserId, assignmentDomain.ScopeId).Scan(&assignmentDomain.UserId, &assignmentDomain.ScopeId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return nil
}

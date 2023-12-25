package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"todolist/helper"
	"todolist/model/domain"
	"todolist/model/dto"
	"todolist/repository"
)

type AssignmentService interface {
	Assign(ctx context.Context, request dto.AssignmentUserScopeRequest) dto.AssignmentUserScopeResponse
}

type AssignmentServiceImpl struct {
	assignmentRepository repository.AssignmentRepository
	db                   *sql.DB
	validate             *validator.Validate
}

func NewAssignmentService(assignmentRepository repository.AssignmentRepository, db *sql.DB, validate *validator.Validate) AssignmentService {
	return &AssignmentServiceImpl{assignmentRepository: assignmentRepository, db: db, validate: validate}
}

func (service *AssignmentServiceImpl) Assign(ctx context.Context, request dto.AssignmentUserScopeRequest) dto.AssignmentUserScopeResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	assignmentUserScopes := domain.AssignmentUserScope{
		UserId:  request.UserId,
		ScopeId: request.ScopeId,
	}

	err = service.assignmentRepository.Find(ctx, tx, assignmentUserScopes)
	if err != nil {
		panic("duplicate assignment found")
	}

	assignmentUserScopeResponse := service.assignmentRepository.Insert(ctx, tx, assignmentUserScopes)

	return helper.AssignmentScopeToResponse(assignmentUserScopeResponse)

}

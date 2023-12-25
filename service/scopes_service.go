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

type ScopeService interface {
	Create(ctx context.Context, request dto.CreateScopeRequest) dto.ScopesResponse
	Delete(ctx context.Context, name string)
}

type ScopesServiceImpl struct {
	scopesRepository repository.ScopesRepository
	db               *sql.DB
	validate         *validator.Validate
}

func NewScopesService(scopesRepository repository.ScopesRepository, db *sql.DB, validate *validator.Validate) ScopeService {
	return &ScopesServiceImpl{scopesRepository: scopesRepository, db: db, validate: validate}
}

func (service *ScopesServiceImpl) Create(ctx context.Context, request dto.CreateScopeRequest) dto.ScopesResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	scopesDomain := domain.ScopesDomain{
		Name: request.Name,
	}

	scopesResponse := service.scopesRepository.Insert(ctx, tx, scopesDomain)

	return helper.ScopesToResponse(scopesResponse)
}

func (service *ScopesServiceImpl) Delete(ctx context.Context, name string) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.scopesRepository.Delete(ctx, tx, name)
}

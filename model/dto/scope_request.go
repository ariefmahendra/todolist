package dto

type CreateScopeRequest struct {
	Name string `json:"name" validate:"required"`
}


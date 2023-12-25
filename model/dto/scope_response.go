package dto

type ScopesResponse struct {
	Name string `json:"name" validate:"required"`
}

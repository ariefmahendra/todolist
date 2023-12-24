package dto

type CreateTodolistRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"isDone"`
}

type UpdateTodolistRequest struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"isDone" validate:"required"`
}

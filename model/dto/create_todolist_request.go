package dto

type CreateTodolistRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"isDone"`
}

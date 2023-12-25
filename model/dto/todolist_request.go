package dto

type CreateTodolistRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"isDone"`
	UserId      int    `json:"user_id" validate:"required"`
}

type UpdateTodolistRequest struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"isDone"`
	UserId      int    `json:"user_id" validate:"required"`
}

type CheckTodolistRequest struct {
	IsDone bool `json:"isDone" validate:"required"`
}

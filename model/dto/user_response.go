package dto

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginUserResponse struct {
	Email  string   `json:"email"`
	Token  string   `json:"token"`
	Scopes []string `json:"scopes"`
}

package dto

type AssignmentUserScopeRequest struct {
	UserId  int `json:"user_id"`
	ScopeId int `json:"scope_id"`
}

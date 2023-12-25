package helper

import (
	"todolist/model/domain"
	"todolist/model/dto"
)

func TodolistToResponse(todolist domain.TodolistDomain) dto.TodolistResponse {
	todolistResponse := dto.TodolistResponse{
		Id:          todolist.Id,
		Title:       todolist.Title,
		Description: todolist.Description,
		IsDone:      todolist.IsDone,
	}
	return todolistResponse
}

func TodolistToResponses(todolist []domain.TodolistDomain) []dto.TodolistResponse {
	var todolistResponses []dto.TodolistResponse
	for _, item := range todolist {
		todoItem := TodolistToResponse(item)
		todolistResponses = append(todolistResponses, todoItem)
	}
	return todolistResponses
}

func UserToResponse(userDomain domain.UserDomain) dto.UserResponse {
	userResponse := dto.UserResponse{
		Id:    userDomain.Id,
		Name:  userDomain.Name,
		Email: userDomain.Email,
	}
	return userResponse
}

func ScopesToResponse(scopesDomain domain.ScopesDomain) dto.ScopesResponse {
	scopesResponse := dto.ScopesResponse{
		Name: scopesDomain.Name,
	}
	return scopesResponse
}

func AssignmentScopeToResponse(assignmentUserScope domain.AssignmentUserScope) dto.AssignmentUserScopeResponse {
	assignmentUserScopeResponse := dto.AssignmentUserScopeResponse{
		UserId:  assignmentUserScope.UserId,
		ScopeId: assignmentUserScope.ScopeId,
	}
	return assignmentUserScopeResponse

}

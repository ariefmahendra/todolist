package helper

import (
	"todolist/model/domain"
	"todolist/model/dto"
)

func ToResponse(todolist domain.TodolistDomain) dto.TodolistResponse {
	todolistResponse := dto.TodolistResponse{
		Id:          todolist.Id,
		Title:       todolist.Title,
		Description: todolist.Description,
		IsDone:      todolist.IsDone,
	}
	return todolistResponse
}

func ToResponses(todolist []domain.TodolistDomain) []dto.TodolistResponse {
	var todolistResponses []dto.TodolistResponse
	for _, item := range todolist {
		todoItem := ToResponse(item)
		todolistResponses = append(todolistResponses, todoItem)
	}
	return todolistResponses
}

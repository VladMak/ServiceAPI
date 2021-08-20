package service

import (
	"github.com/VladMak/ServiceAPI"
	"github.com/VladMak/ServiceAPI/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list ServiceAPI.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
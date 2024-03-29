package service

import (
	"github.com/VladMak/ServiceAPI"
	"github.com/VladMak/ServiceAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user ServiceAPI.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list ServiceAPI.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

package service

import (
	"github.com/VladMak/ServiceAPI/pkg/repository"
	"github.com/VladMak/ServiceAPI"
)

type Authorization interface {
	CreateUser(user ServiceAPI.User) (int, error)
}

type TodoList interface {

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
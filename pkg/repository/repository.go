package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/VladMak/ServiceAPI"
)

type Authorization interface {
	CreateUser(user ServiceAPI.User) (int, error)
	GetUser(username, password string) (ServiceAPI.User, error)
}

type TodoList interface {

}

type TodoItem interface {

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
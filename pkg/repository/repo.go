package repository

import (
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GetUser(username, password string) (todoapp.User, error)
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

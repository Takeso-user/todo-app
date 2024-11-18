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
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(id int) ([]todoapp.TodoList, error)
	GetById(userId int, listId int) (todoapp.TodoList, error)
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
		TodoList:      NewTodoListPostgres(db),
	}
}

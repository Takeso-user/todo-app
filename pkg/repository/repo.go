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
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todoapp.UpdateListInput) error
}
type TodoItem interface {
	Create(id int, item todoapp.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todoapp.TodoItem, error)
	GetById(userId, itemId int) (todoapp.TodoItem, error)
	Delete(userId, itemId int) error
	Update(id int, id2 int, input todoapp.UpdateItemInput) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}

package service

import (
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(userId int) ([]todoapp.TodoList, error)
	GetById(userId int, listId int) (todoapp.TodoList, error)
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
		TodoList:      NewTodoListService(repos.TodoList),
	}

}

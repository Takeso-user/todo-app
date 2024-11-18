package service

import (
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func (t *TodoListService) GetById(userId int, listId int) (todoapp.TodoList, error) {
	return t.repo.GetById(userId, listId)
}

func (t *TodoListService) GetAll(userId int) ([]todoapp.TodoList, error) {
	return t.repo.GetAll(userId)
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (t *TodoListService) Create(userId int, list todoapp.TodoList) (int, error) {
	return t.repo.Create(userId, list)

}

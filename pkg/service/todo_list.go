package service

import (
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/repository"
	"github.com/sirupsen/logrus"
)

type TodoListService struct {
	repo repository.TodoList
}

func (t *TodoListService) Update(userId int, listId int, input todoapp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		logrus.Info("nothing to update. all fields are empty")
		return err
	}
	return t.repo.Update(userId, listId, input)
}

func (t *TodoListService) Delete(userId, id int) error {
	return t.repo.Delete(userId, id)
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

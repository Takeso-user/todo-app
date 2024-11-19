package service

import (
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func (s *TodoItemService) Update(id int, id2 int, input todoapp.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, id2, input)
}

func (s *TodoItemService) Delete(id int, id2 int) error {
	return s.repo.Delete(id, id2)
}

func (s *TodoItemService) GetById(userId, itemId int) (todoapp.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todoapp.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}
func (s TodoItemService) GetAll(UserId, ItemId int) ([]todoapp.TodoItem, error) {

	return s.repo.GetAll(UserId, ItemId)

}

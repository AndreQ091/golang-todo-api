package service

import (
	todo "github.com/AndreQ091/golang-todo"
	"github.com/AndreQ091/golang-todo/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo, listRepo}
}

func (s *TodoItemService) Create(listId int, userId int, input todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(listId, userId)

	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, input)
}

func (s *TodoItemService) GetAll(listId int, userId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(listId, userId)
}

func (s *TodoItemService) GetById(itemId int, userId int) (todo.TodoItem, error) {
	return s.repo.GetById(itemId, userId)
}

func (s *TodoItemService) UpdateById(itemId int, userId int, input todo.UpdateItemInput) error {
	return s.repo.UpdateById(itemId, userId, input)
}

func (s *TodoItemService) DeleteById(itemId int, userId int) error {
	return s.repo.DeleteById(itemId, userId)
}

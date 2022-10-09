package service

import (
	todo "github.com/AndreQ091/golang-todo"
	"github.com/AndreQ091/golang-todo/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(id, userId int) (todo.TodoList, error) {
	return s.repo.GetById(id, userId)
}

func (s *TodoListService) UpdateById(id int, userId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateById(id, userId, input)
}

func (s *TodoListService) DeleteById(id, userId int) error {
	return s.repo.DeleteById(id, userId)
}

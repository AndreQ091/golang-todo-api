package service

import (
	todo "github.com/AndreQ091/golang-todo"
	"github.com/AndreQ091/golang-todo/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(id, userId int) (todo.TodoList, error)
	UpdateById(id int, userId int, input todo.UpdateListInput) error
	DeleteById(id, userId int) error
}

type TodoItem interface {
	Create(listId int, userId int, list todo.TodoItem) (int, error)
	GetAll(listId int, userId int) ([]todo.TodoItem, error)
	GetById(itemId int, userId int) (todo.TodoItem, error)
	UpdateById(id int, userId int, input todo.UpdateItemInput) error
	DeleteById(itemId int, userId int) error
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
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}

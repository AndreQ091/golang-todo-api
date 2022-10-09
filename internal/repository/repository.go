package repository

import (
	todo "github.com/AndreQ091/golang-todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(id, userId int) (todo.TodoList, error)
	UpdateById(id int, userId int, input todo.UpdateListInput) error
	DeleteById(id, userId int) error
}

type TodoItem interface {
	Create(listId int, list todo.TodoItem) (int, error)
	GetAll(listId int, userId int) ([]todo.TodoItem, error)
	GetById(itemId int, userId int) (todo.TodoItem, error)
	UpdateById(id int, userId int, input todo.UpdateItemInput) error
	DeleteById(itemId int, userId int) error
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

package service

import (
	"github.com/ivasdtmbb/todo_list_project/pkg/repository"
	"github.com/ivasdtmbb/todo_list_project/internal/todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
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
		TodoList:	   NewTodoListService(repos.TodoList),
	}
}


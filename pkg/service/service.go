package service

import (
	"github.com/ivasdtmbb/todo_list_project/pkg/repository"
	"github.com/ivasdtmbb/todo_list_project/internal/todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {

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
	}
}


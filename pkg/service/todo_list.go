package service

import (
	//"github.com/ivasdtmbb/todo_list_project"
	"github.com/ivasdtmbb/todo_list_project/internal/todo"
	"github.com/ivasdtmbb/todo_list_project/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
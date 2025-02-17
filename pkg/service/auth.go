package service

import (
	"crypto/sha1"
	_"errors"
	"fmt"

	"github.com/ivasdtmbb/todo_list_project/pkg/repository"
	"github.com/ivasdtmbb/todo_list_project/internal/todo"
)

const (
	salt 		= "lkasjf348759ulkfjl923r8uo"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
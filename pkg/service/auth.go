package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/mohito22/todo-app-golang/models"
	"github.com/mohito22/todo-app-golang/pkg/repository"
)

const sault = "ajaragudju"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(sault)))
}

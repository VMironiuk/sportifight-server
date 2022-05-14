package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/VMironiuk/sportifight-server"
	"github.com/VMironiuk/sportifight-server/pkg/repository"
)

const (
	salt = "hasdfjhsajflb2312bfjlhds232ew3e32fe"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user sportifight.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username string, pawwsord string) (string, error) {
	// TODO: add implementation
	return "test_token", nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", string([]byte(salt)))
}

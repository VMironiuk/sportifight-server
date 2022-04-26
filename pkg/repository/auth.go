package repository

import (
	"github.com/VMironiuk/sportifight-server"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (*AuthRepository) CreateUser(user sportifight.User) (int, error) {
	return 42, nil
}

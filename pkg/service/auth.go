package service

import "github.com/VMironiuk/sportifight-server"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(sportifight.User) (int, error) {
	return 42, nil
}

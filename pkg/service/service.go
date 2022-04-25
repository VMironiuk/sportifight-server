package service

import "github.com/VMironiuk/sportifight-server"

type Auth interface {
	CreateUser(user sportifight.User) (int, error)
}

type Service struct {
	Auth
}

func New() *Service {
	return &Service{
		Auth: NewAuthService(),
	}
}

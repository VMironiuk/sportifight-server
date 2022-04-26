package service

import (
	"github.com/VMironiuk/sportifight-server"
	"github.com/VMironiuk/sportifight-server/pkg/repository"
)

type Auth interface {
	CreateUser(user sportifight.User) (int, error)
}

type Service struct {
	Auth
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}

package repository

import "github.com/VMironiuk/sportifight-server"

type Auth interface {
	CreateUser(user sportifight.User) (int, error)
}

type Repository struct {
	Auth
}

func New() *Repository {
	return &Repository{
		Auth: NewAuthRepository(),
	}
}

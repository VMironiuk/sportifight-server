package repository

import (
	"github.com/VMironiuk/sportifight-server"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user sportifight.User) (int, error)
}

type Repository struct {
	Auth
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}

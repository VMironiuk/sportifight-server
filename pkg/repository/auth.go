package repository

import (
	"github.com/VMironiuk/sportifight-server"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (*AuthRepository) CreateUser(user sportifight.User) (int, error) {
	return 42, nil
}

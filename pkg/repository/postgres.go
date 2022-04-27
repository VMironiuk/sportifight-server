package repository

import "github.com/jmoiron/sqlx"

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=qwerty sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

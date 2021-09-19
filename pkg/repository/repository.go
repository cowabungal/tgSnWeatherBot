package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthRepository(db), NewUserRepository(db)}
}

type Authorization interface {
	IsUser(userId int) error
	CreateUser(username string, userId int) error
}

type User interface {
	Name(userId int) (string, error)
}

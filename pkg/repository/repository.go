package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthRepository(db)}
}

type Authorization interface {
	IsUser(userId int) error
	CreateUser(username string, userId int) error
}

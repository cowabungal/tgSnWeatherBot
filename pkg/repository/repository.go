package repository

import (
	"github.com/jmoiron/sqlx"
	"tgSnWeatherBot"
)

type Repository struct {
	Authorization
	User
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewAuthRepository(db), NewUserRepository(db), NewAdminRepository(db)}
}

type Authorization interface {
	IsUser(userId int) error
	CreateUser(username string, userId int) error
	IsAdmin(userId int) error
	CreateAdmin(username string, userId int) error
}

type User interface {
	Name(userId int) (string, error)
	City(userId int) (string, error)
	ChangeCity (userId int, newCity string) (string, error)
	Info (userId int) (*tgSnWeatherBot.User, error)
	DeleteName(userId int, name string) error
	AddName(userId int, name string) (string, error)
}

type Admin interface {
	UsersList() ([]tgSnWeatherBot.User, error)
}

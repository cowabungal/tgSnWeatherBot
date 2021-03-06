package service

import (
	"tgSnWeatherBot"
	"tgSnWeatherBot/pkg/repository"
)

type Service struct {
	Weather
	Authorization
	User
	Admin
}

func NewService(repo *repository.Repository) *Service {
	return &Service{NewWeatherService(), NewAuthService(repo), NewUserService(repo), NewAdminService(repo)}
}

type Authorization interface {
	IsUser(userId int) error
	CreateUser(username string, userId int) error
	IsAdmin(userId int) error
	CreateAdmin(username string, userId int) error
}

type Weather interface {
	Get(city string) (*tgSnWeatherBot.WeatherData, error)
}

type User interface {
	Name (userId int) (string, error)
	City (userId int) (string, error)
	ChangeCity (userId int, newCity string) (string, error)
	State(userId int) (string, error)
	ChangeState(userId int, newState string) (string, error)
	AddCallbackId(userId int, callbackId string) error
	AddCallbackData(callbackId, callbackData string) error
	GetCallbackData(userId int) (string, error)
	GetCallbackId(userId int) (int, error)
	DeleteCallback(userId int) error
	Info (userId int) (*tgSnWeatherBot.User, error)
	DeleteName(userId int, name string) error
	AddName(userId int, name string) (string, error)
}

type Admin interface {
	UsersList() ([]tgSnWeatherBot.User, error)
}

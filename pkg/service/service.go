package service

import (
	"tgSnWeatherBot"
	"tgSnWeatherBot/pkg/repository"
)

type Service struct {
	Weather
	Authorization
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{NewWeatherService(), NewAuthService(repo), NewUserService(repo)}
}

type Authorization interface {
	IsUser(userId int) error
	CreateUser(username string, userId int) error
}

type Weather interface {
	Get(city string) (*tgSnWeatherBot.WeatherData, error)
}

type User interface {
	Name (userId int) (string, error)
	City (userId int) (string, error)
	ChangeCity (userId int, newCity string) (string, error)
}

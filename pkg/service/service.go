package service

import (
	"tgSnWeatherBot"
	"tgSnWeatherBot/pkg/repository"
)

type Service struct {
	Weather
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{NewWeatherService(), NewAuthService(repo)}
}

type Authorization interface {
	IsUser(userId int) error
	CreateUser(username string, userId int) error
}

type Weather interface {
	Get() (*tgSnWeatherBot.WeatherData, error)
}

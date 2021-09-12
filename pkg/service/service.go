package service

import "tgSnWeatherBot"

type Service struct {
	Weather
}

func NewService() *Service {
	return &Service{Weather: NewWeatherService()}
}

type Weather interface {
	Get() (*tgSnWeatherBot.WeatherData, error)
}

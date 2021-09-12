package service

import (
	owm "github.com/briandowns/openweathermap"
	"os"
	"tgSnWeatherBot"
)

type WeatherService struct {

}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) Get() (*tgSnWeatherBot.WeatherData, error) {
	apiKey := os.Getenv("OWM_API_KEY")

	w, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		return nil, err
	}

	err = w.CurrentByName("Moscow")
	if err != nil {
		return nil, err
	}

	return &tgSnWeatherBot.WeatherData{Temperature: w.Main.Temp}, nil
}

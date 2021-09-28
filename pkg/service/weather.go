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

func (s *WeatherService) Get(city string) (*tgSnWeatherBot.WeatherData, error) {
	apiKey := os.Getenv("OWM_API_KEY")

	w, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		return nil, err
	}

	err = w.CurrentByName(city)
	if err != nil {
		return nil, err
	}

	return &tgSnWeatherBot.WeatherData{Temperature: w.Main.Temp}, nil
}

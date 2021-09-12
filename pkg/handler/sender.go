package handler

import (
	"fmt"
	"os"
	"tgSnWeatherBot"
)

func weatherMessage(data *tgSnWeatherBot.WeatherData) string {
	return fmt.Sprintf(os.Getenv("SUN_NAME") + ", температура в Москве: %.0f°C", data.Temperature)
}

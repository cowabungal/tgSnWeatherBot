package handler

import (
	"fmt"
	"math/rand"
	"os"
	"tgSnWeatherBot"
	"time"
)

var Names = make([]string, 0, 0)

func weatherMessage(data *tgSnWeatherBot.WeatherData) string {
	return fmt.Sprintf(pickName(Names) + ", температура в Москве: %.0f°C", data.Temperature)
}

func initNames(names []string) []string {
	names = append(names, os.Getenv("NAME1"))
	names = append(names, os.Getenv("NAME2"))
	names = append(names, os.Getenv("NAME3"))

	return names
}

func pickName(names []string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(names))

	return names[randomIndex]
}

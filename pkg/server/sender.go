package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"tgSnWeatherBot"
)

func weatherMessage(data *tgSnWeatherBot.WeatherData, name string) string {
	return fmt.Sprintf("%s, температура в Москве: %.0f°C", name, data.Temperature)
}

func (s *Server) GetUserName(userId int) string {
	name, err := s.service.User.Name(userId)
	if err != nil {
		logrus.Error("error: server: GetUserName: " + err.Error())
		return "Юзер"
	}

	return name
}

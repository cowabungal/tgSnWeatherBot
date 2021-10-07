package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"tgSnWeatherBot"
)

func weatherMessage(data *tgSnWeatherBot.WeatherData, name string, user *tgSnWeatherBot.User) string {
	return fmt.Sprintf("%s, температура в городе %s: %.0f°C", name, user.City, data.Temperature)
}

func profileMessage(user *tgSnWeatherBot.User) string {
	return fmt.Sprintf(
		"Username: %s\n" +
			"UserId: %d\n" +
			"Город: %s\n\n",
			user.Username, user.UserId, user.City)
}

func (s *Server) GetUserName(userId int) string {
	name, err := s.service.User.Name(userId)
	if err != nil {
		logrus.Error("error: server: GetUserName: " + err.Error())
		return "Юзер"
	}

	return name
}

func sendingSettingMsg(user *tgSnWeatherBot.User) string {
	return fmt.Sprintf("")
}

func textMessage(name string) string {
	return fmt.Sprintf("%s, нажми на кнопку 'Погода'", name)
}

func usersListMessage(list []tgSnWeatherBot.User) string {
	return fmt.Sprintf("Количество пользователей: %d", len(list))
}

func userSettingsMessage(user *tgSnWeatherBot.User) string {
	return fmt.Sprintf("Настройки пользователя: %s", user.Username)
}

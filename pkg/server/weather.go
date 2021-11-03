package server

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func (s *Server) GetWeather (m *telebot.Message) {
	logrus.Printf("message from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	user := s.NewUser(m.Sender)
	s.service.ChangeState(m.Sender.ID, "default")

	weatherData, err := s.service.Weather.Get(user.City)
	if err != nil {
		logrus.Error("getWeather: Get: " + err.Error())
		return
	}

	name := s.GetUserName(m.Sender.ID)

	_, err = s.bot.Send(m.Sender, weatherMessage(weatherData, name, user))
	if err != nil {
		logrus.Error("getWeather: Send: " + err.Error())
		return
	}

	logrus.Printf("Bot send message: 'weatherMessage' to %s", m.Sender.Username)
}

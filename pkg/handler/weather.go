package handler

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func (h *Handler) getWeather (m *telebot.Message) {
	logrus.Printf("message from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	weatherData, err := h.service.Weather.Get()
	if err != nil {
		logrus.Error("getWeather: Get: " + err.Error())
		return
	}

	_, err = h.bot.Send(m.Sender, weatherMessage(weatherData))
	if err != nil {
		logrus.Error("getWeather: Send: " + err.Error())
		return
	}

	logrus.Printf("Bot send message: <<weatherMessage>> to %s", m.Sender.Username)
}

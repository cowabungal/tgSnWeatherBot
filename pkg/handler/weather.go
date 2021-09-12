package handler

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func (h *Handler) getWeather (m *telebot.Message) {
	weatherData, err := h.service.Weather.Get()
	if err != nil {
		logrus.Error("getWeather: Get: " + err.Error())
		return
	}

	h.bot.Send(m.Sender, weatherMessage(weatherData))
}

package handler

import (
	"gopkg.in/tucnak/telebot.v2"
	"tgSnWeatherBot/pkg/service"
)

type Handler struct {
	service *service.Service
	bot *telebot.Bot
}

func (h *Handler) InitRoutes() {
	Names = initNames(Names)

	h.bot.Handle(telebot.OnText, h.getWeather)
}

func NewHandler(s *service.Service, b *telebot.Bot) *Handler {
	return &Handler{service: s, bot: b}
}

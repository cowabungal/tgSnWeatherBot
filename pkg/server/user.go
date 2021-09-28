package server

import (
	"gopkg.in/tucnak/telebot.v2"
	"tgSnWeatherBot"
)

func (s *Server) NewUser(u *telebot.User) *tgSnWeatherBot.User {
	city, err := s.service.User.City(u.ID)
	if err != nil {
		city = "Москва"
	}

	return &tgSnWeatherBot.User{
		Username: u.Username,
		UserId:   u.ID,
		Names:    nil,
		City:     city,
	}
}

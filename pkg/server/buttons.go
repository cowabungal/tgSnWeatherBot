package server

import (
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"tgSnWeatherBot"
)

type Buttons struct {
	Button *telebot.ReplyMarkup
}

func NewButtons(b *telebot.ReplyMarkup) *Buttons {
	return &Buttons{Button: b}
}

func (s *Buttons) Main() *telebot.ReplyMarkup {
	main := s.Button
	getWeather := main.Text("Погода")
	profile := main.Text("Профиль")
	main.Reply(
		main.Row(getWeather, profile),
		)

	return main
}

func (s *Buttons) ProfileInline() *telebot.ReplyMarkup {
	profile := s.Button
	changeCity := changeCity(profile)
	profile.Inline(
		profile.Row(changeCity),
	)

	return profile
}

func changeCity(b *telebot.ReplyMarkup) telebot.Btn {
	// значение unique обязательно должно быть цельным
	return b.Data("Изменить город", "city")
}

func (s *Buttons) Settings(user *tgSnWeatherBot.User) *telebot.ReplyMarkup {
	settings := s.Button

	username := settings.Text(fmt.Sprintf("Username: %s", user.Username))
	userId := settings.Text(fmt.Sprintf("UserId: %d", user.UserId))
	city := settings.Text(fmt.Sprintf("Город: %s", user.City))

	settings.Reply(
		settings.Row(username),
		settings.Row(userId),
		settings.Row(city),
		)

	return settings
}

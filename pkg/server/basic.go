package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func (s *Server) mainButtons (m *telebot.Message) {
	logrus.Printf("password from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	s.bot.Send(m.Sender, "Привет. Я - бот, который подскажет тебе погоду на улице.", s.buttons.Main())
}

func (s *Server) profile (m *telebot.Message) {
	logrus.Printf("profile from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	user := s.NewUser(m.Sender)
	profileInline := s.buttons.ProfileInline()
	s.bot.Send(m.Sender, profileMessage(user), profileInline)
}

func (s *Server) changeCity (c *telebot.Callback) {
	err := s.bot.Respond(c, &telebot.CallbackResponse{Text: "Отправь название города"})
	if err != nil {
		logrus.Error("changeCity: Respond: " + err.Error())
	}

	logrus.Printf("city from: %s; id: %d; ms: %s", c.Sender.Username, c.Sender.ID, c.Data)

	s.bot.Send(c.Sender, "Отправь название города")
	s.bot.Handle(telebot.OnText, s.changeCityAns)
}

func (s *Server) changeCityAns (m *telebot.Message) {
	logrus.Printf("city from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	user := s.NewUser(m.Sender)
	city := m.Text

	city, err := s.service.User.ChangeCity(user.UserId, city)
	if err != nil {
		s.bot.Send(m.Sender, "Ошибка в названии города.")
		return
	}

	s.bot.Send(m.Sender, fmt.Sprintf("Город успешно изменен на: %s", city))
}


func (s *Server) text (m *telebot.Message) {
	logrus.Printf("message from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	name := s.GetUserName(m.Sender.ID)
	s.bot.Send(m.Sender, textMessage(name), s.buttons.Main())
}

func (s *Server) settingsBut (m *telebot.Message) {
	user := s.NewUser(m.Sender)
	s.bot.Send(m.Sender, "Settings!", s.buttons.Settings(user))
}
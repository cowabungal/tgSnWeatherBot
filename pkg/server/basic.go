package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func (s *Server) mainButtons (m *telebot.Message) {
	logrus.Printf("password from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	s.bot.Send(m.Sender, "Привет. Я - бот, который подскажет тебе погоду на улице.", s.button.Main())
}

func (s *Server) profile (m *telebot.Message) {
	logrus.Printf("profile from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	user := s.NewUser(m.Sender)
	profileInline, cityBut := s.button.ProfileInline()
	s.bot.Send(m.Sender, profileMessage(user), profileInline)
	s.bot.Handle(&cityBut, s.changeCity)
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

func (s *Server) changeCityAdm (c *telebot.Callback) {
	err := s.bot.Respond(c, &telebot.CallbackResponse{Text: "Отправь название города"})
	if err != nil {
		logrus.Error("changeCity: Respond: " + err.Error())
	}

	logrus.Printf("city from: %s; id: %d; ms: %s", c.Sender.Username, c.Sender.ID, c.Data)

	s.bot.Send(c.Sender, "Отправь название города")

	s.data.prevCallback = c
	s.bot.Handle(telebot.OnText, s.changeCityAnsAdm)
}

func (s *Server) changeCityAnsAdm (m *telebot.Message) {
	logrus.Printf("city from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	c := s.data.prevCallback

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("citySettings: getUser: " + err.Error())
		return
	}

	city := m.Text

	city, err = s.service.User.ChangeCity(user.UserId, city)
	if err != nil {
		s.bot.Send(m.Sender, "Ошибка в названии города.")
		return
	}

	s.bot.Send(m.Sender, fmt.Sprintf("Город успешно изменен на: %s", city))
}

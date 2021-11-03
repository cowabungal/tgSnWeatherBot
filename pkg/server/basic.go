package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func (s *Server) mainButtons(m *telebot.Message) {
	logrus.Printf("password from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	main := s.button.Main()
	s.bot.Send(m.Sender, "Привет. Я - бот, который подскажет тебе погоду на улице.", &main)
}

func (s *Server) profile(m *telebot.Message) {
	logrus.Printf("profile from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	user, err := s.getUser(m.Sender.ID)
	if err != nil {
		logrus.Error("profile: getUser: " + err.Error())
		return
	}

	main, cityBut, namesBut := s.button.UserSettings(user)
	//profileInline, cityBut, sendingSetting := s.button.ProfileInline()
	s.bot.Send(m.Sender, profileMessage(user), &main)

	s.bot.Handle(&cityBut, s.citySettings)
	s.bot.Handle(&namesBut, s.namesSettings)
}

func (s *Server) changeCity(c *telebot.Callback) {
	err := s.bot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		logrus.Error("changeCity: Respond: " + err.Error())
	}

	logrus.Printf("city from: %s; id: %d; ms: %s", c.Sender.Username, c.Sender.ID, c.Data)

	s.bot.Send(c.Sender, "Отправь название города")
	s.service.User.ChangeState(c.Sender.ID, "changeCity")

	s.data.prevCallback = c
	err = s.service.User.AddCallbackId(c.Sender.ID, c.ID)
	if err != nil {
		logrus.Error("changeCity: AddCallbackId: " + err.Error())
	}

	err = s.service.User.AddCallbackData(c.ID, c.Data)
	if err != nil {
		logrus.Error("changeCity: AddCallbackId: " + err.Error())
	}
}

func (s *Server) changeCityAns (m *telebot.Message) {
	_, err := s.service.User.GetCallbackId(m.Sender.ID)
	if err != nil {
		logrus.Error("changeCityAns: GetCallbackId: " + err.Error())
		return
	}

	logrus.Printf("city from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	user := s.NewUser(m.Sender)
	city := m.Text

	city, err = s.service.User.ChangeCity(user.UserId, city)
	if err != nil {
		s.bot.Send(m.Sender, "Ошибка в названии города.")
		return
	}

	s.bot.Send(m.Sender, fmt.Sprintf("Город успешно изменен на: %s", city))
	s.service.User.ChangeState(user.UserId, "default")
}

func (s *Server) changeCityAdm (c *telebot.Callback) {
	err := s.bot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		logrus.Error("changeCitAdmy: Respond: " + err.Error())
	}

	logrus.Printf("city from: %s; id: %d; ms: %s", c.Sender.Username, c.Sender.ID, c.Data)

	s.bot.Send(c.Sender, "Отправь название города")
	s.service.User.ChangeState(c.Sender.ID, "changeCityAdm")

	err = s.service.User.AddCallbackId(c.Sender.ID, c.ID)
	if err != nil {
		logrus.Error("changeCityAdm: AddCallbackId: " + err.Error())
	}
	err = s.service.User.AddCallbackData(c.ID, c.Data)
	if err != nil {
		logrus.Error("changeCityAdm: AddCallbackData: " + err.Error())
	}
}

func (s *Server) changeCityAnsAdm (m *telebot.Message) {
	userId, err := s.service.User.GetCallbackData(m.Sender.ID)
	if err != nil {
		logrus.Error("changeCityAnsAdm: GetCallbackData: " + err.Error())
		return
	}

	logrus.Printf("city from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	userIdInt, _ := strconv.Atoi(userId)
	user, err := s.getUser(userIdInt)
	if err != nil {
		logrus.Error("citySettingsAnsAdm: getUser: " + err.Error())
		return
	}

	city := m.Text

	city, err = s.service.User.ChangeCity(user.UserId, city)
	if err != nil {
		s.bot.Send(m.Sender, "Ошибка в названии города.")
		return
	}

	s.bot.Send(m.Sender, fmt.Sprintf("Город успешно изменен на: %s", city))
	s.service.User.ChangeState(m.Sender.ID, "default")
	s.service.User.DeleteCallback(m.Sender.ID)
}

func notOwner(c *telebot.Callback, m *telebot.Message) bool {
	if c.Sender.ID != m.Sender.ID {
		return true
	}

	return false
}

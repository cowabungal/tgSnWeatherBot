package server

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func (s *Server) text (m *telebot.Message) {
	logrus.Printf("message from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	name := s.GetUserName(m.Sender.ID)
	state := s.GetUserState(m.Sender.ID)

	switch state {
	case "default":
		s.bot.Send(m.Sender, textMessage(name))
	case "changeCity":
		s.changeCityAns(m)
	case "changeCityAdm":
		s.changeCityAnsAdm(m)
	case "addName":
		s.addName(m)
	case "resendMessage":
		s.resendMessage(m)
	case "adminPass":
		s.adminPass(m)
	}
}

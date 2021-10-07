package server

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func (s *Server) text (m *telebot.Message) {
	logrus.Printf("message from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	name := s.GetUserName(m.Sender.ID)
	s.bot.Send(m.Sender, textMessage(name))
}

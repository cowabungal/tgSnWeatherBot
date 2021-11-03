package server

import (
	"gopkg.in/tucnak/telebot.v2"
	"os"
	"tgSnWeatherBot/pkg/service"
)

type AuthServer struct {
	service *service.Service
}

func NewAuthServer(s *service.Service) *AuthServer {
	return &AuthServer{service: s}
}

func (s *AuthServer) isUser(userId int) bool {
	err := s.service.Authorization.IsUser(userId)
	if err != nil {
		return false
	}

	return true
}

func isPassword(password string) bool {
	return password == os.Getenv("BOT_PASSWORD")
}

func (s *AuthServer) createUser(username string, userId int) {
	err := s.service.Authorization.CreateUser(username, userId)
	if err != nil {
		return
	}
}

func (s *Server) adminPass (m *telebot.Message) {
	if os.Getenv("ADMIN_PASSWORD") == m.Text {
		err := s.service.Authorization.CreateAdmin(m.Sender.Username, m.Sender.ID)
		if err != nil {
			return
		}

		adminBut := s.button.MainAdmin()
		s.bot.Send(m.Sender, "Вы успешно залогинены в аккаунт администратора.", adminBut)
	} else {
		s.bot.Send(m.Sender, "Пароль неверный. Нажмите /admin и введите корректный пароль для доступа к админ-панели.")
	}

	s.service.User.ChangeState(m.Sender.ID, "default")
}

package server

import (
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

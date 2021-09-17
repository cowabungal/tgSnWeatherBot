package server

import (
	"tgSnWeatherBot/pkg/service"
)

type AuthServer struct {
	service *service.Service
}

func NewAuthServer(service *service.Service) *AuthServer {
	return &AuthServer{service: service}
}

func (s *AuthServer) isUser(userId int) bool {
	err := s.service.Authorization.IsUser(userId)
	if err != nil {
		return false
	}

	return true
}

func isPassword(password string) bool {
	return password == "123"
}

func (s *AuthServer) createUser(username string, userId int) {
	err := s.service.Authorization.CreateUser(username, userId)
	if err != nil {
		return
	}
}

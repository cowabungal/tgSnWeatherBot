package service

import "tgSnWeatherBot/pkg/repository"

// AuthService структура сервиса авторизации
type AuthService struct {
	repo *repository.Repository
}

// NewAuthService возвращает указатель на новую структуру AuthService
func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

// IsUser принимает id пользователя и проверяет есть ли он в базе данных
func (s *AuthService) IsUser(userId int) error {
	return s.repo.Authorization.IsUser(userId)
}

func (s *AuthService) CreateUser(username string, userId int) error {
	return s.repo.Authorization.CreateUser(username, userId)
}

func (s *AuthService) IsAdmin(userId int) error {
	return s.repo.Authorization.IsAdmin(userId)
}

func (s *AuthService) CreateAdmin(username string, userId int) error {
	return s.repo.Authorization.CreateAdmin(username, userId)
}

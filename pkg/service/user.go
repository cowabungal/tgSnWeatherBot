package service

import "tgSnWeatherBot/pkg/repository"

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Name(userId int) (string, error) {
	return s.repo.User.Name(userId)
}

package service

import (
	"tgSnWeatherBot"
	"tgSnWeatherBot/pkg/repository"
)

type AdminService struct {
	repo *repository.Repository
}

func NewAdminService(repo *repository.Repository) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) UsersList() ([]tgSnWeatherBot.User, error) {
	return s.repo.Admin.UsersList()
}

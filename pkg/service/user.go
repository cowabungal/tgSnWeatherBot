package service

import (
	"github.com/sirupsen/logrus"
	"tgSnWeatherBot/pkg/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Name(userId int) (string, error) {
	return s.repo.User.Name(userId)
}

func (s *UserService) City(userId int) (string, error) {
	return s.repo.User.City(userId)
}

func (s *UserService) ChangeCity(userId int, newCity string) (string, error) {
	err := validateCity(newCity)

	if err != nil {
		return "", err
	}

	return s.repo.User.ChangeCity(userId, newCity)
}

func validateCity(newCity string) error {
	ws := NewWeatherService()
	_, err := ws.Get(newCity)
	if err != nil {
		logrus.Error("service: validateCity: Get: " + err.Error())
		return err
	}

	return nil
}

package service

import (
	"github.com/sirupsen/logrus"
	"tgSnWeatherBot"
	"tgSnWeatherBot/pkg/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddName(userId int, name string) (string, error) {
	return s.repo.User.AddName(userId, name)
}

func (s *UserService) Name(userId int) (string, error) {
	return s.repo.User.Name(userId)
}

func (s *UserService) DeleteName(userId int, name string) error {
	return s.repo.User.DeleteName(userId, name)
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

func (s *UserService) State(userId int) (string, error) {
	return s.repo.User.State(userId)
}

func (s *UserService) ChangeState(userId int, newState string) (string, error) {
	return s.repo.User.ChangeState(userId, newState)
}

func (s *UserService) AddCallbackId(userId int, callbackId string) error {
	return s.repo.User.AddCallbackId(userId, callbackId)
}

func (s *UserService) AddCallbackData(callbackId, callbackData string) error {
	return s.repo.User.AddCallbackData(callbackId, callbackData)
}

func (s *UserService) GetCallbackData(userId int) (string, error) {
	return s.repo.User.GetCallbackData(userId)
}

func (s *UserService) GetCallbackId(userId int) (int, error) {
	return s.repo.User.GetCallbackId(userId)
}

func (s *UserService) DeleteCallback(userId int) error {
	return s.repo.User.DeleteCallback(userId)
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

func (s *UserService) Info(userId int) (*tgSnWeatherBot.User, error) {
	return s.repo.User.Info(userId)
}

package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"tgSnWeatherBot/pkg/handler"
	"tgSnWeatherBot/pkg/service"
)

func main() {
	// загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	srv := handler.NewBotServer()
	services := service.NewService()
	handlers := handler.NewHandler(services, srv.Bot)

	srv.Run(handlers)
}

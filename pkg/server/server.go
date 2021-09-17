package server

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"tgSnWeatherBot/pkg/service"
	"time"
)

type Server struct {
	service *service.Service
	bot *telebot.Bot
}

func (s *Server) InitRoutes() {
	Names = initNames(Names)

	s.bot.Handle(telebot.OnText, s.getWeather)
}

func (s *Server) Run() {
	s.InitRoutes()
	logrus.Info("The BotServer has successfully run")
	s.bot.Start()
}

func NewBotServer(s *service.Service) *Server {
	authServer := NewAuthServer(s)

	poller := &telebot.LongPoller{Timeout: 15 * time.Second}

	authMiddleware := telebot.NewMiddlewarePoller(poller, func(upd *telebot.Update) bool {
		if authServer.isUser(upd.Message.Sender.ID) {
			return true
		}

		if isPassword(upd.Message.Text) {
			authServer.createUser(upd.Message.Sender.Username, upd.Message.Sender.ID)
			return true
		}

		bot, _ := telebot.NewBot(telebot.Settings{
			Token:  os.Getenv("BOT_TOKEN"),
			Poller: poller,
		})

		bot.Send(upd.Message.Sender, "Введи пароль, чтобы продолжить работу с ботом.")

		return false
	})

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: authMiddleware,
	})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Server{service: s, bot: bot}
}

package handler

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

type Server struct {
	Bot *telebot.Bot
}

func (s *Server) Run(h *Handler) {
	h.InitRoutes()
	logrus.Info("The BotServer has successfully run")
	s.Bot.Start()
}

func NewBotServer() *Server {
	bot, err := telebot.NewBot(telebot.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "https://api.telegram.org",

		Token:  "1442509990:AAE4L0qMEhebu1xHNWSnqTBx7o5SQGEOtlI",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Server{Bot: bot}
}

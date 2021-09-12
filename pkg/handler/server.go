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
	poller := &telebot.LongPoller{Timeout: 15 * time.Second}
	authMiddleware := telebot.NewMiddlewarePoller(poller, func(upd *telebot.Update) bool {
		if !isUser(upd.Message.Sender.Username) {
			return false
		}

		return true
	})

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  "1442509990:AAE4L0qMEhebu1xHNWSnqTBx7o5SQGEOtlI",
		Poller: authMiddleware,
	})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Server{Bot: bot}
}

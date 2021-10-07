package server

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"tgSnWeatherBot/pkg/server/buttons"
	"tgSnWeatherBot/pkg/service"
	"time"
)

type Server struct {
	service    *service.Service
	bot    *telebot.Bot
	button *buttons.Buttons
	data   *Data
}

type Data struct {
	prevCallback *telebot.Callback
}

func NewData() *Data {
	return &Data{}
}

func (s *Server) InitRoutes() {
	s.bot.Handle("Погода", s.GetWeather)
	s.bot.Handle("Профиль", s.profile)
	s.bot.Handle("Юзер-панель", s.mainButtons)
	s.bot.Handle("Пользователи", s.usersList)
	s.bot.Handle(os.Getenv("BOT_PASSWORD"), s.mainButtons)
	s.bot.Handle("/start", s.mainButtons)
	s.bot.Handle("/admin", s.admin)
	s.bot.Handle(telebot.OnText, s.text)
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

		if upd.Message == nil {
			switch upd.Query {
			case nil:
				if upd.Callback == nil {
					return false
				} else {
					if authServer.isUser(upd.Callback.Sender.ID) {
						return true
					}
				}
			default:
				if authServer.isUser(upd.Query.From.ID) {
					return true
				}
			}
		} else {
			if authServer.isUser(upd.Message.Sender.ID) {
				return true
			}

			if isPassword(upd.Message.Text) {
				authServer.createUser(upd.Message.Sender.Username, upd.Message.Sender.ID)
				return true
			}
		}

		bot, _ := telebot.NewBot(telebot.Settings{
			Token:  os.Getenv("BOT_TOKEN"),
			Poller: poller,
		})

		logrus.Printf("message from: %s; id: %d; ms: %s", upd.Message.Sender.Username, upd.Message.Sender.ID, upd.Message.Text)
		bot.Send(upd.Message.Sender, "Введи пароль, чтобы продолжить работу с ботом.")

		logrus.Printf("Bot send message: <<weatherMessage>> to %s", upd.Message.Sender.Username)

		return false
	})

	b, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: authMiddleware,
	})

	if err != nil {
		log.Fatal(err)
		return nil
	}
	menu := telebot.ReplyMarkup{ResizeReplyKeyboard: true}
	bu := buttons.NewButtons(menu)
	data := NewData()
	return &Server{service: s, bot: b, button: bu, data: data}
}

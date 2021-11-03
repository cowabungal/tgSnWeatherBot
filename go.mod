module tgSnWeatherBot

// +heroku goVersion go1.16
go 1.16

require (
	github.com/briandowns/openweathermap v0.16.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	gopkg.in/tucnak/telebot.v2 v2.4.0
	gopkg.in/tucnak/telebot.v3 v3.0.0-20211015201320-13d54ae7338e
)

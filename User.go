package tgSnWeatherBot

type User struct {
	Username string `db:"username"`
	UserId int `db:"user_id"`
	Names []string `db:"names"`
	City string `db:"city"`
}

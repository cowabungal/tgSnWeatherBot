package tgSnWeatherBot

import "strconv"

type User struct {
	Username string `db:"username"`
	UserId int `db:"user_id"`
	Names []string `db:"names"`
	City string `db:"city"`
	SendingTime string `db:"sending_time"`
}

func (u *User) Recipient() string {
	return strconv.Itoa(u.UserId)
}

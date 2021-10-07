package buttons

import (
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"tgSnWeatherBot"
)

type Buttons struct {
	Button *telebot.ReplyMarkup
}

func NewButtons(b *telebot.ReplyMarkup) *Buttons {
	return &Buttons{Button: b}
}

func (s *Buttons) Main() *telebot.ReplyMarkup {
	main := s.Button
	getWeather := main.Text("Погода")
	profile := main.Text("Профиль")
	main.Reply(
		main.Row(getWeather, profile),
		)

	return main
}

func (s *Buttons) MainAdmin() *telebot.ReplyMarkup {
	main := s.Button

	usersList := main.Text("Пользователи")
	user := main.Text("Юзер-панель")
	main.Reply(
		main.Row(usersList),
		main.Row(user),
	)

	return main
}

func changeCityBut(b *telebot.ReplyMarkup) telebot.Btn {
	return b.Data("Изменить город", "city")
}

func sendingSettingBut(b *telebot.ReplyMarkup) telebot.Btn {
	return b.Data("Настроить рассылку", "sending")
}

func changeCityAdmBut(user *tgSnWeatherBot.User, b *telebot.ReplyMarkup) telebot.Btn {
	return b.Data("Изменить город", "city", strconv.Itoa(user.UserId))
}

func usernameBut(i int, user tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(user.Username, fmt.Sprintf("username%d", i), strconv.Itoa(user.UserId))
}

func cityBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(fmt.Sprintf("Город: %s", user.City), "citySettings", strconv.Itoa(user.UserId))
}

func namesCountBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(fmt.Sprintf("Имена: %dшт.", len(user.Names)), "nameSettings", strconv.Itoa(user.UserId))
}

func returnBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data("🔙 Назад", "return", strconv.Itoa(user.UserId))
}

func addNameBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data("✅ Добавить имя", "addName", strconv.Itoa(user.UserId))
}

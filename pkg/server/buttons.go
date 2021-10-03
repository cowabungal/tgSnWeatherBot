package server

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
	admin := main.Text("Админ-панель")
	main.Reply(
		main.Row(getWeather, profile),
		main.Row(admin),
		)

	return main
}

func (s *Buttons) ProfileInline() *telebot.ReplyMarkup {
	profile := s.Button
	changeCity := changeCity(profile)
	profile.Inline(
		profile.Row(changeCity),
	)

	return profile
}

func changeCity(b *telebot.ReplyMarkup) telebot.Btn {
	// значение unique обязательно должно быть цельным
	return b.Data("Изменить город", "city")
}

func changeCityAdm(user *tgSnWeatherBot.User, b *telebot.ReplyMarkup) telebot.Btn {
	return b.Data("Изменить город", "city", strconv.Itoa(user.UserId))
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


func (s *Buttons) UserListInline(list []tgSnWeatherBot.User) (*telebot.ReplyMarkup, []telebot.Btn) {
	main := s.Button

	inlineList := userInline(list, main)

	main.Inline(
		inlineList,
	)
	return main, inlineList
}

func userInline (list []tgSnWeatherBot.User, main *telebot.ReplyMarkup) []telebot.Btn {
	usernameInlineList := make([]telebot.Btn, 0)

	for i := 0; i < len(list); i++ {
		usernameInlineList = append(usernameInlineList, usernameInline(i, list[i], main))
	}

	return usernameInlineList
}

func usernameInline(i int, user tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(user.Username, fmt.Sprintf("username%d", i), strconv.Itoa(user.UserId))
}

// userSettings return main, city, names
func (s *Buttons) userSettings(user *tgSnWeatherBot.User) (*telebot.ReplyMarkup, telebot.Btn, telebot.Btn) {
	main := s.Button

	city := cityBut(user, main)
	names := namesBut(user, main)

	main.Inline(
		main.Row(city),
		main.Row(names),
		)

	return main, city, names
}

func cityBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(fmt.Sprintf("Город: %s", user.City), "citySettings", strconv.Itoa(user.UserId))
}

func namesBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(fmt.Sprintf("Имена: %dшт.", len(user.Names)), "nameSettings", strconv.Itoa(user.UserId))
}

func returnBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data("🔙 Назад", "return", strconv.Itoa(user.UserId))
}

func (s *Buttons) returnInline(user *tgSnWeatherBot.User) (*telebot.ReplyMarkup, telebot.Btn) {
		main := s.Button
		returnBut := returnBut(user, main)

		main.Inline(
			main.Row(returnBut),
			)

		return main, returnBut
}

func addNameBut(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data("✅ Добавить имя", "addName", strconv.Itoa(user.UserId))
}

// citySettings return main, city, changeCity, returnBut
func (s *Buttons) citySettings(user *tgSnWeatherBot.User) (*telebot.ReplyMarkup, telebot.Btn, telebot.Btn, telebot.Btn) {
	main := s.Button

	city := cityBut(user, main)
	changeCity := changeCityAdm(user, main)
	returnBut := returnBut(user, main)

	main.Inline(
		main.Row(city),
		main.Row(changeCity),
		main.Row(returnBut),
	)

	return main, city, changeCity, returnBut
}

func (s *Buttons) NamesListInline(user *tgSnWeatherBot.User) (*telebot.ReplyMarkup, telebot.Btn, telebot.Btn) {
	main := s.Button

	inlineList, returnBut, nameAddBut := namesInline(user, main)

	main.InlineKeyboard = inlineList
	return main, returnBut, nameAddBut
}

func namesInline (user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) ([][]telebot.InlineButton, telebot.Btn, telebot.Btn) {
	inlineList := make([][]telebot.InlineButton, 0)
	inlineButton := make([]telebot.InlineButton, 0)
	list := user.Names

	for i := 0; i < len(list); i++ {
		btn := nameInlineSlc(i, user, main)
		b := btn[0]
		bi := b.Inline()

		inlineButton = append(inlineButton, *bi)
		inlineList = append(inlineList, inlineButton)
		inlineButton = make([]telebot.InlineButton, 0)
	}

	addName := addNameBut(user, main)
	b := addName.Inline()
	inlineButton = append(inlineButton, *b)
	inlineList = append(inlineList, inlineButton)
	inlineButton = make([]telebot.InlineButton, 0)

	returnBut := returnBut(user, main)
	b = returnBut.Inline()
	inlineButton = append(inlineButton, *b)
	inlineList = append(inlineList, inlineButton)
	inlineButton = make([]telebot.InlineButton, 0)

	return inlineList, returnBut, addName
}

func nameInlineSlc(i int, user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) []telebot.Btn {
	btn := make([]telebot.Btn, 0)
	btn = append(btn, main.Data(user.Names[i], fmt.Sprintf("name%d", i), strconv.Itoa(user.UserId), user.Names[i]))

	return btn
}

// NameInline return main, nameIn, delete, returnBut
func (s *Buttons) NameInline(name string, userId int) (*telebot.ReplyMarkup, telebot.Btn, telebot.Btn, telebot.Btn) {
	main := s.Button

	nameIn := main.Data(name, "nameInline",  strconv.Itoa(userId), name)
	deleteBut := main.Data("Удалить имя", "deleteName",  strconv.Itoa(userId), name)
	returnBut := returnBut(&tgSnWeatherBot.User{UserId: userId}, main)

	main.Inline(
		main.Row(nameIn),
		main.Row(deleteBut),
		main.Row(returnBut),
	)
	return main, nameIn, deleteBut, returnBut
}

// YesOrNo return main, yes, no
func (s *Buttons) YesOrNo(name string, userId int) (*telebot.ReplyMarkup, telebot.Btn, telebot.Btn) {
	main := s.Button

	yes := main.Data("Да", "Yes",  strconv.Itoa(userId), name)
	no := main.Data("Нет", "No",  strconv.Itoa(userId), name)

	main.Inline(
		main.Row(yes, no),
	)
	return main, yes, no
}

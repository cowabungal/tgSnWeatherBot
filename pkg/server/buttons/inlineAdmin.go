package buttons

import (
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"tgSnWeatherBot"
)

// UserList возвращает инлайн кнопку со списком пользователей по username
func (s *Buttons) UserList(list []tgSnWeatherBot.User) (telebot.ReplyMarkup, []telebot.Btn) {
	main := s.Button

	butList := userSlc(list, &main)

	main.Inline(
		butList,
	)
	return main, butList
}

// userSlc принимает список пользователей и возвращает слайс кнопок с username
func userSlc(list []tgSnWeatherBot.User, main *telebot.ReplyMarkup) []telebot.Btn {
	usernameList := make([]telebot.Btn, 0)

	for i, v := range list {
		usernameList = append(usernameList, usernameBut(i, v, main))
	}

	return usernameList
}

// UserSettings return main, city, namesCount
func (s *Buttons) UserSettings(user *tgSnWeatherBot.User) (telebot.ReplyMarkup, telebot.Btn, telebot.Btn) {
	main := s.Button

	city := cityBut(user, &main)
	namesCount := namesCountBut(user, &main)

	main.Inline(
		main.Row(city),
		main.Row(namesCount),

	)

	return main, city, namesCount
}

func (s *Buttons) ReturnInline(user *tgSnWeatherBot.User) (telebot.ReplyMarkup, telebot.Btn) {
	main := s.Button
	returnBut := returnBut(user, &main)

	main.Inline(
		main.Row(returnBut),
	)

	return main, returnBut
}

// CitySettings return main, city, changeCityBut, returnBut
func (s *Buttons) CitySettings(user *tgSnWeatherBot.User) (telebot.ReplyMarkup, telebot.Btn, telebot.Btn, telebot.Btn) {
	main := s.Button

	city := cityBut(user, &main)
	changeCity := changeCityAdmBut(user, &main)
	returnBut := returnBut(user, &main)

	main.Inline(
		main.Row(city),
		main.Row(changeCity),
		main.Row(returnBut),
	)

	return main, city, changeCity, returnBut
}

func (s *Buttons) NamesList(user *tgSnWeatherBot.User) (telebot.ReplyMarkup, telebot.Btn, telebot.Btn) {
	main := s.Button

	inlineList, returnBut, nameAddBut := names(user, &main)

	main.InlineKeyboard = inlineList
	return main, returnBut, nameAddBut
}

func names(user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) ([][]telebot.InlineButton, telebot.Btn, telebot.Btn) {
	inlineList := make([][]telebot.InlineButton, 0)
	inlineButton := make([]telebot.InlineButton, 0)
	list := user.Names

	for i := 0; i < len(list); i++ {
		btn := nameSlc(i, user, main)
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

func nameSlc(i int, user *tgSnWeatherBot.User, main *telebot.ReplyMarkup) []telebot.Btn {
	btn := make([]telebot.Btn, 0)
	btn = append(btn, main.Data(user.Names[i], fmt.Sprintf("name%d", i), strconv.Itoa(user.UserId), user.Names[i]))

	return btn
}

// Name return main, nameIn, delete, returnBut
func (s *Buttons) Name(name string, userId int) (telebot.ReplyMarkup, telebot.Btn, telebot.Btn, telebot.Btn) {
	main := s.Button

	nameIn := main.Data(name, "nameInline",  strconv.Itoa(userId), name)
	deleteBut := main.Data("Удалить имя", "deleteName",  strconv.Itoa(userId), name)
	returnBut := returnBut(&tgSnWeatherBot.User{UserId: userId}, &main)

	main.Inline(
		main.Row(nameIn),
		main.Row(deleteBut),
		main.Row(returnBut),
	)
	return main, nameIn, deleteBut, returnBut
}

// YesOrNo return main, yes, no
func (s *Buttons) YesOrNo(name string, userId int) (telebot.ReplyMarkup, telebot.Btn, telebot.Btn) {
	main := s.Button

	yes := main.Data("Да", "Yes",  strconv.Itoa(userId), name)
	no := main.Data("Нет", "No",  strconv.Itoa(userId), name)

	main.Inline(
		main.Row(yes, no),
	)
	return main, yes, no
}

package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"strings"
)

func (s *Server) admin (m *telebot.Message) {
	logrus.Printf("admin from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	err := s.service.Authorization.IsAdmin(m.Sender.ID)
	if err != nil {
		s.bot.Send(m.Sender, "Введите пароль для доступа к админ-панели.")
		s.service.User.ChangeState(m.Sender.ID, "adminPass")
	} else {
		adminBut := s.button.MainAdmin()
		s.bot.Send(m.Sender, "Вы успешно вошли в аккаунт администратора.", &adminBut)
	}
}

func (s *Server) usersList (m *telebot.Message) {
	logrus.Printf("usersList from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	err := s.service.Authorization.IsAdmin(m.Sender.ID)
	if err != nil {
		return
	}

	// получаем список юзеров
	usersList, err := s.service.Admin.UsersList()
	if err != nil {
		return
	}

	// получаем сгусток инлайн кнопок и массив самих кнопок
	usersListInline, usersButtons := s.button.UserList(usersList)
	s.bot.Send(m.Sender, usersListMessage(usersList), &usersListInline)

	// обработчик нажатий на юзера
	for _, v := range usersButtons {
		s.bot.Handle(&v, s.userSettings)
	}
}

func (s *Server) userSettings(c *telebot.Callback) {
	logrus.Printf("inlineUser from: %s; id: %d; data: %s", c.Sender.Username, c.Sender.ID, c.Data)
	s.bot.Respond(c, &telebot.CallbackResponse{})

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("inlineUser: getUser: " + err.Error())
		return
	}

	s.service.ChangeState(c.Sender.ID, "default")

	main, cityBut, namesBut := s.button.UserSettings(user)

	s.bot.Edit(c.Message, userSettingsMessage(user), &main)

	s.bot.Handle(&cityBut, s.citySettings)
	s.bot.Handle(&namesBut, s.namesSettings)
}

func (s *Server) citySettings(c *telebot.Callback) {
	logrus.Printf("citySettings from: %s; id: %d; data: %s", c.Sender.Username, c.Sender.ID, c.Data)
	s.bot.Respond(c, &telebot.CallbackResponse{})

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("citySettings: getUser: " + err.Error())
		return
	}

	main, city, changeCity, returnBut := s.button.CitySettings(user)

	s.bot.Edit(c.Message, userSettingsMessage(user), &main)

	s.bot.Handle(&city, s.citySettings)
	s.bot.Handle(&changeCity, s.changeCityAdm)
	s.bot.Handle(&returnBut, s.userSettings)
}

func (s *Server) namesSettings(c *telebot.Callback) {
	logrus.Printf("namesSettings from: %s; id: %d; data: %s", c.Sender.Username, c.Sender.ID, c.Data)
	s.bot.Respond(c, &telebot.CallbackResponse{})

	s.service.User.ChangeState(c.Sender.ID, "default")

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("nameSettings: getUser: " + err.Error())
		return
	}

	// получаем сгусток инлайн кнопок и массив самих кнопок
	namesListInline, returnBut, nameAddBut := s.button.NamesList(user)
	s.bot.Edit(c.Message, userSettingsMessage(user), &namesListInline)

	list := namesListInline.InlineKeyboard[:len(namesListInline.InlineKeyboard)]

	for _, v := range list {
		for _, va := range v {
			s.bot.Handle(&va, s.nameSettings)
		}
	}

	s.bot.Handle(&returnBut, s.userSettings)
	s.bot.Handle(&nameAddBut, s.preAddName)
}

func (s *Server) nameSettings(c *telebot.Callback) {
	logrus.Printf("nameSettings from: %s; id: %d; data: %s", c.Sender.Username, c.Sender.ID, c.Data)
	s.bot.Respond(c, &telebot.CallbackResponse{})

	data := strings.Split(c.Data, "|")
	userId, _ := strconv.Atoi(data[0])
	name := data[1]

	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("nameSettings: getUser: " + err.Error())
		return
	}

	main, nameIn, deleteBut, returnBut := s.button.Name(name, userId)

	s.bot.Edit(c.Message, userSettingsMessage(user), &main)
	s.bot.Handle(&nameIn, s.nameSettings)
	s.bot.Handle(&deleteBut, s.preDeleteName)
	s.bot.Handle(&returnBut, s.namesSettings)
}

func (s *Server) preDeleteName(c *telebot.Callback) {
	logrus.Printf("preDeleteName from: %s; id: %d; data: %s", c.Sender.Username, c.Sender.ID, c.Data)
	s.bot.Respond(c, &telebot.CallbackResponse{})

	data := strings.Split(c.Data, "|")
	userId, _ := strconv.Atoi(data[0])
	name := data[1]

	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("preDelete: getUser: " + err.Error())
		return
	}

	main, yes, no := s.button.YesOrNo(name, userId)

	s.bot.Edit(c.Message,fmt.Sprintf("Вы уверены, что хотите удалить имя '%s' у пользователя: %s?", name, user.Username), &main)
	s.bot.Handle(&yes, s.deleteNameAdm)
	s.bot.Handle(&no, s.nameSettings)
}

func (s *Server) deleteNameAdm(c *telebot.Callback) {
	logrus.Printf("preDeleteName from: %s; id: %d; data: %s", c.Sender.Username, c.Sender.ID, c.Data)
	s.bot.Respond(c, &telebot.CallbackResponse{})

	data := strings.Split(c.Data, "|")
	userId, _ := strconv.Atoi(data[0])
	name := data[1]

	err := s.service.User.DeleteName(userId, name)
	if err != nil {
		logrus.Error("deleteNameAdm: DeleteName: " + err.Error())
		return
	}

	s.bot.Send(c.Sender, "Имя успешно удалено!")
}

func (s *Server) preAddName(c *telebot.Callback) {
	logrus.Printf("preAddName from: %s; id: %d; data: %s", c.Sender.Username, c.Sender.ID, c.Data)
	s.bot.Respond(c, &telebot.CallbackResponse{})

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("preAddName: getUser: " + err.Error())
		return
	}

	main, returnBut := s.button.ReturnInline(user)

	s.bot.Edit(c.Message,fmt.Sprintf("Отправьте новое имя для пользователя: %s", user.Username), &main)
	s.service.User.ChangeState(c.Sender.ID, "addName")
	s.bot.Handle(&returnBut, s.namesSettings)

	s.data.prevCallback = c
}

func (s *Server) addName(m *telebot.Message) {
	c := s.data.prevCallback
	//if notOwner(c, m) {
	//	return
	//}

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("preAddName: getUser: " + err.Error())
		return
	}

	name, err := s.service.User.AddName(userId, m.Text)

	s.bot.Send(m.Sender, fmt.Sprintf("Имя '%s' для пользователя: %s успешно добавлено.", name, user.Username))
	s.service.User.ChangeState(m.Sender.ID, "default")
}

func (s *Server) usersListMessage(m *telebot.Message) {
	logrus.Printf("usersListMessage from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	err := s.service.Authorization.IsAdmin(m.Sender.ID)
	if err != nil {
		return
	}

	// получаем список юзеров
	usersList, err := s.service.Admin.UsersList()
	if err != nil {
		return
	}

	// получаем сгусток инлайн кнопок и массив самих кнопок
	usersListInline, usersButtons := s.button.UserList(usersList)
	s.bot.Send(m.Sender, usersListMessage(usersList), &usersListInline)

	// обработчик нажатий на юзера
	for _, v := range usersButtons {
		s.bot.Handle(&v, s.sendMessage)
	}
}

func (s *Server) sendMessage(c *telebot.Callback) {
	err := s.bot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		logrus.Error("sendMessage: Respond: " + err.Error())
	}

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("sendMessage: getUser: " + err.Error())
		return
	}

	logrus.Printf("sendMessage from: %s; id: %d; ms: %s", c.Sender.Username, c.Sender.ID, c.Data)

	main, cancelBut := s.button.CancelInline(user)

	s.bot.Edit(c.Message, fmt.Sprintf("Отправка сообщения пользователю: %s", user.Username), &main)
	s.bot.Send(c.Sender, "Отправь текст сообщения")
	s.service.ChangeState(c.Sender.ID, "resendMessage")

	s.data.prevCallback = c
	s.bot.Handle(&cancelBut, s.resendMessageCancel)
}

func (s *Server) resendMessage(m *telebot.Message) {
	c := s.data.prevCallback
	if notOwner(c, m) {
		return
	}

	logrus.Printf("resendMessage from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	userId, _ := strconv.Atoi(c.Data)
	user, err := s.getUser(userId)
	if err != nil {
		logrus.Error("resendMessage: getUser: " + err.Error())
		return
	}

	_, err = s.bot.Send(user, m.Text)
	if err != nil {
		logrus.Error("resendMessage: bot.Send: " + err.Error())
		s.bot.Send(m.Sender, "Ошибка. Сообщение не доставлено.")
		return
	}

	s.bot.Send(c.Sender, "Сообщение успешно доставлено")
}

func (s *Server) resendMessageCancel(c *telebot.Callback) {
	err := s.bot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		logrus.Error("resendMessageCancel: Respond: " + err.Error())
	}

	s.bot.Send(c.Sender,"Отправка сообщения отменена")
	s.service.User.ChangeState(c.Sender.ID, "default")
}

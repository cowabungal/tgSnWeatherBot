package buttons

import "gopkg.in/tucnak/telebot.v2"

func (s *Buttons) ProfileInline() (*telebot.ReplyMarkup, telebot.Btn) {
	profile := s.Button
	changeCity := changeCityBut(profile)
	profile.Inline(
		profile.Row(changeCity),
	)

	return profile, changeCity
}

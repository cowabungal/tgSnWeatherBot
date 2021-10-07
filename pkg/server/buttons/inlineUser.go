package buttons

import (
	"gopkg.in/tucnak/telebot.v2"
)

func (s *Buttons) ProfileInline() (telebot.ReplyMarkup, telebot.Btn, telebot.Btn) {
	profile := s.Button
	changeCity := changeCityBut(&profile)
	sendingSetting := sendingSettingBut(&profile)

	profile.Inline(
		profile.Row(changeCity, sendingSetting),
	)

	return profile, changeCity, sendingSetting
}

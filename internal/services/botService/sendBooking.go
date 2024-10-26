package botService

import (
	"fmt"
	t "time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bs *botService) SendBooking(
	name, time, date, phone string,
	quantity int,
	is_confirmed bool,
) error {
	var confirmed string

	if is_confirmed {
		confirmed = "Да"
	} else {
		confirmed = "Нет"
	}

	parsedDate, errParse := t.Parse(t.RFC3339, date)
	if errParse != nil {
		return errParse
	}

	text := fmt.Sprintf(
		"<b>Новая заявка на бронирование: </b>\nИмя: %s\nКол-во человек: %d\nПодтверждена: %s\nВремя: %s\nДата: %s\nНом. телефона: <a href='tel:%s'>%s</a>",
		name, quantity, confirmed, time, parsedDate.Format("02/01/2006"), phone, phone,
	)
	msg := tgbotapi.NewMessage(bs.chatId, text)
	msg.ParseMode = "HTML"

	_, errSend := bs.bot.Send(msg)
	return errSend
}

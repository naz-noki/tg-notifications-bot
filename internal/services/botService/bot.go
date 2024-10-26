package botService

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type botService struct {
	bot    *tgbotapi.BotAPI
	chatId int64
	token  string
}

func New(
	token string,
	chatId int64,
) (*botService, error) {
	bs := new(botService)

	bot, errNewBotAPI := tgbotapi.NewBotAPI(token)
	if errNewBotAPI != nil {
		return nil, errNewBotAPI
	}

	bs.chatId = chatId
	bs.token = token
	bs.bot = bot

	return bs, nil
}

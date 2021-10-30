package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
)

func (c *Commander) CheckCommand(chatID int64, userID int, args string) {
	_ = userID
	_ = args

	msg := tgbotapi.NewMessage(chatID, answer.NotImplemented)
	c.bot.Send(msg)
}

package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

func (c *Commander) HelpCommand(userID model.UserID, tg tgUserInfo, args string) {
	_ = userID
	_ = args

	msg := tgbotapi.NewMessage(tg.ChatID, answer.Help)
	c.bot.Send(msg)
}

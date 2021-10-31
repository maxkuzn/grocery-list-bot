package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
)

func (c *Commander) HelpCommand(tg tgMeta, args string) {
	_ = args

	msg := tgbotapi.NewMessage(tg.ChatID, answer.Help)
	c.bot.Send(msg)
}

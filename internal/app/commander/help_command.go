package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) HelpCommand(chatID int, userID int, args string) {
	_ = userID
	_ = args

	msg := tgbotapi.NewMessage(chatID, answer.Help)
	r.bot.Send(msg)
}

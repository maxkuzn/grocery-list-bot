package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot: bot,
	}
}

func (c *Commander) HandleMessage(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID

	if !update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.OnlyCommands)
		c.bot.Send(msg)
		return
	}

	args := update.Message.CommandArguments()
	switch update.Message.Command() {
	case "help":
		c.Help(chatID, userID, args)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.UnknownCommand)
		c.bot.Send(msg)
		return
	}
}

func (c *Commander) HandleCallback(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.NotImplemented)
	c.bot.Send(msg)
}

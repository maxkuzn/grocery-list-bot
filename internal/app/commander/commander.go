package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
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
		c.HelpCommand(chatID, userID, args)
	case "create":
		c.CreateCommand(chatID, userID, args)
	case "switch":
		c.SwitchCommand(chatID, userID, args)
	case "delete":
		c.DeleteCommand(chatID, userID, args)
	case "add":
		c.AddCommand(chatID, userID, args)
	case "remove":
		c.RemoveCommand(chatID, userID, args)
	case "check":
		c.CheckCommand(chatID, userID, args)
	case "uncheck":
		c.UncheckCommand(chatID, userID, args)
	case "show":
		c.ShowCommand(chatID, userID, args)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.UnknownCommands)
		c.bot.Send(msg)
		return
	}
}

func (c *Commander) HandleCallback(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.NotImplemented)
	c.bot.Send(msg)
}

package commander

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/app/idbinder"
	"github.com/maxkuzn/grocery-list-bot/internal/service/listsdb"
)

type tgMeta struct {
	ChatID   int64
	UserID   int
	UserName string
}

type Commander struct {
	bot      *tgbotapi.BotAPI
	db       listsdb.ListsDB
	idBinder *idbinder.IDBinder
}

func New(bot *tgbotapi.BotAPI, db listsdb.ListsDB) *Commander {
	return &Commander{
		bot:      bot,
		db:       db,
		idBinder: idbinder.New(),
	}
}

func (c *Commander) HandleMessage(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID
	userName := update.Message.From.UserName
	info := tgMeta{
		ChatID:   chatID,
		UserID:   userID,
		UserName: userName,
	}

	if !update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.OnlyCommands)
		c.bot.Send(msg)
		return
	}

	args := strings.TrimSpace(update.Message.CommandArguments())
	switch update.Message.Command() {
	case "help":
		c.HelpCommand(info, args)
	case "create":
		c.CreateCommand(info, args)
	case "switch":
		c.SwitchCommand(chatID, userID, args)
	case "delete":
		c.DeleteCommand(info, args)
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

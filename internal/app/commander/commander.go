package commander

import (
	"errors"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/maxkuzn/grocery-list-bot/internal/app/sender"
	"github.com/maxkuzn/grocery-list-bot/internal/service/usersdb"
)

type tgUserInfo struct {
	ChatID   int64
	UserID   int64
	UserName string
}

type Commander struct {
	usersDB usersdb.UsersDB

	sender *sender.Sender
}

func New(bot *tgbotapi.BotAPI, usersDB usersdb.UsersDB, s *sender.Sender) *Commander {
	return &Commander{
		usersDB: usersDB,
		sender:  s,
	}
}

func (c *Commander) HandleMessage(update tgbotapi.Update) {
	info := tgUserInfo{
		ChatID:   update.Message.Chat.ID,
		UserID:   update.Message.From.ID,
		UserName: update.Message.From.UserName,
	}

	userID, err := c.usersDB.GetUserID(info.UserID)
	if errors.Is(err, usersdb.ErrNotFound) {
		userID, err = c.usersDB.CreateUser(info.UserID)
	}

	if err != nil {
		c.sender.SendInternalError(info.ChatID, err)
		return
	}

	if !update.Message.IsCommand() {
		// TODO: accept two messaged commands
		c.sender.SendOnlyCommands(info.ChatID)
		return
	}

	args := strings.TrimSpace(update.Message.CommandArguments())

	switch update.Message.Command() {
	case "help":
		c.HelpCommand(userID, info, args)
	case "create":
		// c.CreateCommand(userID, info, args)
		c.sender.SendNotImplemented(update.Message.Chat.ID)
	case "add":
		// c.AddCommand(userID, info, args)
		c.sender.SendNotImplemented(update.Message.Chat.ID)
	case "show":
		// c.ShowCommand(userID, info, args)
		c.sender.SendNotImplemented(update.Message.Chat.ID)
	default:
		c.sender.SendUnknownCommand(info.ChatID, update.Message.Command())

		return
	}
}

func (c *Commander) HandleCallback(update tgbotapi.Update) {
	c.sender.SendNotImplemented(update.Message.Chat.ID)
}

package commander

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/app/idbinder"
	"github.com/maxkuzn/grocery-list-bot/internal/app/metainfo"
	"github.com/maxkuzn/grocery-list-bot/internal/service/listsdb"
)

type tgUserInfo struct {
	ChatID   int64
	UserID   int
	UserName string
}

type Commander struct {
	bot      *tgbotapi.BotAPI
	db       listsdb.ListsDB
	idBinder *idbinder.IDBinder
	metaInfo *metainfo.MetaInfoStorer
}

func New(bot *tgbotapi.BotAPI, db listsdb.ListsDB) *Commander {
	return &Commander{
		bot:      bot,
		db:       db,
		idBinder: idbinder.New(),
		metaInfo: metainfo.New(),
	}
}

func (c *Commander) HandleMessage(update tgbotapi.Update) {
	info := tgUserInfo{
		ChatID:   update.Message.Chat.ID,
		UserID:   update.Message.From.ID,
		UserName: update.Message.From.UserName,
	}

	userID, ok := c.idBinder.Tg2User(info.UserID)
	if !ok {
		userID = c.db.CreateUser()
		err := c.idBinder.BindUser(info.UserID, userID)
		if err != nil {
			c.send(info.ChatID, answer.InternalError(err))
			return
		}
	}

	if !update.Message.IsCommand() {
		// TODO: accept two messaged commands
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.OnlyCommands)
		c.bot.Send(msg)
		return
	}

	args := strings.TrimSpace(update.Message.CommandArguments())
	switch update.Message.Command() {
	case "help":
		c.HelpCommand(userID, info, args)
	case "create":
		c.CreateCommand(userID, info, args)
	case "switch":
		c.SwitchCommand(userID, info, args)
	case "delete":
		c.DeleteCommand(userID, info, args)
	case "add":
		c.AddCommand(userID, info, args)
	case "remove":
		c.RemoveCommand(userID, info, args)
	case "check":
		c.CheckCommand(userID, info, args)
	case "uncheck":
		c.UncheckCommand(userID, info, args)
	case "show":
		c.ShowCommand(userID, info, args)
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

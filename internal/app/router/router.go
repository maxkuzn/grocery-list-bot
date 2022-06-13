package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/maxkuzn/grocery-list-bot/internal/app/commander"
	"github.com/maxkuzn/grocery-list-bot/internal/app/sender"
)

type Router struct {
	commander *commander.Commander
	sender    *sender.Sender
}

func New(commander *commander.Commander, sender *sender.Sender) *Router {
	return &Router{
		commander: commander,
		sender:    sender,
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	if appropriateMessage(update) {
		r.commander.HandleMessage(update)
		return
	}

	if update.CallbackQuery != nil {
		r.commander.HandleCallback(update)
		return
	}

	// TODO(sender): use constants
	r.sender.SendText(update.Message.Chat.ID, "cannot respond")
}

func appropriateMessage(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}

	if update.Message.Chat == nil {
		return false
	}

	if update.Message.From == nil {
		return false
	}

	if len(update.Message.Text) == 0 {
		return false
	}

	if len(update.Message.From.UserName) == 0 {
		return false
	}

	return true
}

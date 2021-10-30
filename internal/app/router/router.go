package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/app/commander"
)

type Router struct {
	bot       *tgbotapi.BotAPI
	commander *commander.Commander
}

func New(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot:       bot,
		commander: commander.New(bot),
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	if appropriateMessage(update.Message) {
		r.commander.HandleMessage(update)
		return
	}

	if update.CallbackQuery != nil {
		r.commander.HandleCallback(update)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer.CannotRespond)
	r.bot.Send(msg)
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
	return true
}

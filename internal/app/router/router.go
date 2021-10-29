package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Router struct {
	bot *tgbotapi.BotApi
}

func New(bot *tgbotapi.BotApi) *Router {
	return Router{
		bot: bot,
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	r.bot.Send(msg)
}

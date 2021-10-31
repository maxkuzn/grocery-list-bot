package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

func (c *Commander) ShowCommand(userID model.UserID, tg tgUserInfo, args string) {
	_ = args

	metaInfo := c.metaInfo.Get(userID)
	if !metaInfo.AnyListSelected {
		c.send(tg.ChatID, answer.ListNotSelected)
		return
	}
	list, err := c.db.GetList(userID, metaInfo.SelectedList)
	if err != nil {
		c.send(tg.ChatID, answer.InternalError(err))
		return
	}

	c.metaInfo.SetList(userID, list)

	msg := tgbotapi.NewMessage(tg.ChatID, answer.ShowList(list))
	c.bot.Send(msg)
}

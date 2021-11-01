package commander

import (
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

	c.send(tg.ChatID, answer.ShowList(list))
}

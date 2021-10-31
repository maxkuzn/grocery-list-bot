package commander

import (
	"strconv"

	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

func (c *Commander) RemoveCommand(userID model.UserID, tg tgUserInfo, args string) {
	itemIndex, err := strconv.Atoi(args)
	if err != nil {
		c.send(tg.ChatID, answer.RemoveHelp)
		return
	}

	metaInfo := c.metaInfo.Get(userID)
	if !metaInfo.AnyListSelected {
		c.send(tg.ChatID, answer.ListNotSelected)
		return
	}
	// len = N
	// 0...N-1
	// 1...N
	if itemIndex > len(metaInfo.Items) {
		c.send(tg.ChatID, answer.InvalidIndex)
		return
	}
	itemID := metaInfo.Items[itemIndex-1].ID

	err = c.db.RemoveItem(userID, metaInfo.SelectedList, itemID)
	if err != nil {
		c.send(tg.ChatID, answer.InternalError(err))
		return
	}

	c.send(tg.ChatID, answer.ItemDeleted())
}

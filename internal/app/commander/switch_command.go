package commander

import (
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

func (c *Commander) SwitchCommand(userID model.UserID, tg tgUserInfo, args string) {
	listName := args
	if len(listName) == 0 {
		c.send(tg.ChatID, answer.SwitchHelp)
		return
	}

	listID, ok := c.idBinder.ListName2ID(tg.UserName, listName)
	if !ok {
		c.send(tg.ChatID, answer.ListDoesntExist(listName))
		return
	}

	c.metaInfo.SelectList(userID, listID)
	c.send(tg.ChatID, answer.ListSelected(listName, listID))
}

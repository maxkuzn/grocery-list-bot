package commander

import (
	"errors"

	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/model"
	"github.com/maxkuzn/grocery-list-bot/internal/service/listsdb"
)

func (c *Commander) DeleteCommand(userID model.UserID, tg tgUserInfo, args string) {
	listName := args
	if len(listName) == 0 {
		c.send(tg.ChatID, answer.DeleteHelp)
		return
	}

	listID, ok := c.idBinder.ListName2ID(tg.UserName, listName)
	if !ok {
		c.send(tg.ChatID, answer.ListDoesntExist(listName))
		return
	}
	err := c.db.RemoveList(userID, listID)
	if err != nil {
		switch {
		case errors.Is(err, listsdb.ListNameDuplicate):
			c.send(tg.ChatID, answer.ListNameExists(listName))
		default:
			c.send(tg.ChatID, answer.InternalError(err))
		}
		return
	}

	c.metaInfo.UnsetList(userID)
	c.send(tg.ChatID, answer.ListDeleted(listName, listID))
}

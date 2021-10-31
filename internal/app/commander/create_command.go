package commander

import (
	"errors"

	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/service/listsdb"
)

func (c *Commander) CreateCommand(tg tgMeta, args string) {
	userID, ok := c.idBinder.Tg2User(tg.UserID)
	if !ok {
		userID = c.db.CreateUser()
		err := c.idBinder.BindUser(tg.UserID, userID)
		if err != nil {
			c.send(tg.ChatID, answer.InternalError(err))
			return
		}
	}

	listName := args
	if len(listName) == 0 {
		c.send(tg.ChatID, answer.CreateHelp)
		return
	}

	listID, err := c.db.CreateList(userID, listName)
	if err != nil {
		switch {
		case errors.Is(err, listsdb.ListNameDuplicate):
			c.send(tg.ChatID, answer.ListNameExists(listName))
		default:
			c.send(tg.ChatID, answer.InternalError(err))
		}
		return
	}

	c.idBinder.BindList(tg.UserName, listName, listID)
	c.send(tg.ChatID, answer.ListCreated(listName, listID))
}

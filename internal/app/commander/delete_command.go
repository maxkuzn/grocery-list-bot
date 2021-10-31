package commander

import (
	"errors"

	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/service/listsdb"
)

func (c *Commander) DeleteCommand(tg tgMeta, args string) {
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
		c.send(tg.ChatID, answer.DeleteHelp)
		return
	}

	listID, ok := c.idBinder.ListName2ID(tg.UserName, listName)
	if !ok {
		c.send(tg.ChatID, answer.ListDoesntExist(listName))
		// TODO: send error
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

	c.send(tg.ChatID, answer.ListCreated(listName, listID))
}

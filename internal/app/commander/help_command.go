package commander

import (
	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

func (c *Commander) HelpCommand(userID model.UserID, tg tgUserInfo, args string) {
	_ = userID
	_ = args

	c.sender.SendText(tg.ChatID, "some help message /help")
}

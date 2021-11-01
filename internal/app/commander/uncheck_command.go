package commander

import (
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
	"github.com/maxkuzn/grocery-list-bot/internal/model"
)

func (c *Commander) UncheckCommand(userID model.UserID, tg tgUserInfo, args string) {
	_ = userID
	_ = args

	c.send(tg.ChatID, answer.NotImplemented)
}

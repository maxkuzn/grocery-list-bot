package commander

/*
func (c *Commander) AddCommand(userID model.UserID, tg tgUserInfo, args string) {
	itemDescription := args
	if len(itemDescription) == 0 {
		c.send(tg.ChatID, answer.AddHelp)
		return
	}

	metaInfo := c.metaInfo.Get(userID)
	if !metaInfo.AnyListSelected {
		c.send(tg.ChatID, answer.ListNotSelected)
		return
	}

	err := c.db.AddItem(userID, metaInfo.SelectedList, model.Item{
		Description: itemDescription,
	})
	if err != nil {
		c.send(tg.ChatID, answer.InternalError(err))
		return
	}

	c.send(tg.ChatID, answer.ItemAdded())
}
*/

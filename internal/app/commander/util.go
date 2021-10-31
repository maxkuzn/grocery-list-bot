package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) send(chatID int64, msgText string) {
	msg := tgbotapi.NewMessage(chatID, msgText)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("send error: %v", err)
	}
}

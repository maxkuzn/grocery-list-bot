package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxkuzn/grocery-list-bot/internal/app/answer"
)

func (c *Commander) send(chatID int64, msgText string) {
	msg := tgbotapi.NewMessage(chatID, answer.InsertMarkdown(msgText))
	msg.ParseMode = "MarkdownV2"
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("send error: %v", err)
	}
}

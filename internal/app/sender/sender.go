package sender

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Sender struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Sender {
	return &Sender{
		bot: bot,
	}
}

func (s *Sender) sendText(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := s.bot.Send(msg)
	if err != nil {
		log.Printf("send error: %v", err)
	}
}

func (s *Sender) SendText(chatID int64, text string) {
	s.sendText(chatID, text)
}

func (s *Sender) SendInternalError(chatID int64, err error) {
	s.sendText(chatID, "internal error:\n"+err.Error())
}

func (s *Sender) SendNotImplemented(chatID int64) {
	s.sendText(chatID, "not implemented")
}

func (s *Sender) SendUnknownCommand(chatID int64, command string) {
	s.sendText(chatID, fmt.Sprintf("unknown command %q", command))
}

func (s *Sender) SendOnlyCommands(chatID int64) {
	s.sendText(chatID, "currently bot accepts only commands")
}

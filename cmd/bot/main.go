package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/maxkuzn/grocery-list-bot/internal/app/commander"
	"github.com/maxkuzn/grocery-list-bot/internal/app/router"
	"github.com/maxkuzn/grocery-list-bot/internal/app/sender"
	"github.com/maxkuzn/grocery-list-bot/internal/service/usersdb/inmemory"
)

func main() {
	// Setup bot
	_ = godotenv.Load()

	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	defer bot.StopReceivingUpdates()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Setup service

	usersDB := inmemory.NewInMemoryUsersDB()

	senderV := sender.New(bot)

	commanderV := commander.New(bot, usersDB, senderV)

	routerHandler := router.New(commanderV, senderV)

	run(bot, routerHandler)
}

func run(bot *tgbotapi.BotAPI, routerHandler *router.Router) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Run bot
	updateCfg := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(updateCfg)
mainLoop:
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				break mainLoop
			}

			routerHandler.HandleUpdate(update)
		case <-ctx.Done():
			break mainLoop
		}
	}
}

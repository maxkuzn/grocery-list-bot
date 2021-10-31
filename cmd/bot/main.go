package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/maxkuzn/grocery-list-bot/internal/app/router"
	"github.com/maxkuzn/grocery-list-bot/internal/service/listsdb"
)

func main() {
	// Setup bot
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Setup service
	db := listsdb.NewInMemoryDB()

	// Run bot
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	routerHandler := router.New(bot, db)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}

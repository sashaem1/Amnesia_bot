package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/sashaem1/Amnesia_bot/pkg/telegram"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	telegramAPIkey, _ := os.LookupEnv("TELEGRAM_API_KEY")

	b, err := telegram.NewTgBot(telegramAPIkey)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

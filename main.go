package main

import (
	"discord_bot/bot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	bot.Run(token)
}

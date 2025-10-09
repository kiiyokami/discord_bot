package main

import (
	"discord_bot/bot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	bot.Run(os.Getenv("DISCORD_TOKEN"))
}

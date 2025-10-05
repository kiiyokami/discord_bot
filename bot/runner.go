package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(token string) {
	discord, err := discordgo.New("Bot " + token)
	checkError(err)

	discord.AddHandler(MessageHandler)

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	err = discord.Open()
	checkError(err)

	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}

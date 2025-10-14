package commands

import (
	"discord_bot/services"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func FetchImage(s *discordgo.Session, m *discordgo.MessageCreate) {
	words := strings.Fields(strings.ToLower(m.Content))
	for _, word := range words {
		path := fmt.Sprintf("public/%s", word)
		if _, err := os.Stat(path); err == nil {
			fetchAndSendImage(s, m, path)
			return
		}
	}
}

func fetchAndSendImage(s *discordgo.Session, m *discordgo.MessageCreate, path string) {
	imagePath, err := services.RandomImageFetcher(path)
	if err != nil || imagePath == "" {
		return
	}
	
	file, err := os.Open(fmt.Sprintf("%s/%s", path, imagePath))
	if err != nil {
		return
	}
	defer file.Close()
	
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Files: []*discordgo.File{{Name: imagePath, Reader: file}},
	})
}

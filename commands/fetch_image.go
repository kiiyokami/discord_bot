package commands

import (
	"discord_bot/services"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func FetchImage(s *discordgo.Session, m *discordgo.MessageCreate) {
	path := fmt.Sprintf("public/%s", strings.ToLower(m.Content))
	if _, err := os.Stat(path); err == nil {
		imagePath, _ := services.RandomImageFetcher(path)
		absPath := fmt.Sprintf("public/%s/%s", strings.ToLower(m.Content), imagePath)
		file, _ := os.Open(absPath)
		defer file.Close()
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Files: []*discordgo.File{{Name: imagePath, Reader: file}},
		})
	}
}

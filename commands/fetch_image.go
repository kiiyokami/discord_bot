package commands

import (
	"discord_bot/services"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, content string, files []*discordgo.File) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content, Files: files},
	})
}

func FetchImage(s *discordgo.Session, i *discordgo.InteractionCreate) {
	path := fmt.Sprintf("public/%s", i.ApplicationCommandData().Options[0].StringValue())
	imagePath, err := services.RandomImageFetcher(path)
	if err != nil {
		respond(s, i, "No image found", nil)
		return
	}
	file, err := os.Open(fmt.Sprintf("%s/%s", path, imagePath))
	if err != nil {
		respond(s, i, "Error opening image", nil)
		return
	}
	defer file.Close()
	respond(s, i, "", []*discordgo.File{{Name: imagePath, Reader: file}})
}

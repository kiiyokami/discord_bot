package commands

import (
	"discord_bot/services"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ConvertTime(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	hour, minute, success := services.TimeConverter(int(opts[0].IntValue()), int(opts[1].IntValue()), opts[2].StringValue())

	message := "❌ Incorrect country. Please check the country name."
	if success {
		message = fmt.Sprintf("🕐 **Base Time:** %02d:%02d → **%s Time:** %02d:%02d", opts[0].IntValue(), opts[1].IntValue(), opts[2].StringValue(), hour, minute)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: message},
	})
}

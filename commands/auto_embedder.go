package commands

import (
	"discord_bot/services"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var domains = []string{"instagram.com", "pixiv.com", "pixiv.net", "x.com", "tiktok.com", "imgur.com", "twitter.com", "reddit.com", "facebook.com"}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func AutoEmbedder(s *discordgo.Session, m *discordgo.MessageCreate) {
	if u, err := url.ParseRequestURI(m.Content); err == nil && contains(domains, strings.TrimPrefix(u.Host, "www.")) {
		webhook, _ := s.WebhookCreate(m.ChannelID, "AutoEmbed", "")
		_, _ = s.WebhookExecute(webhook.ID, webhook.Token, true, &discordgo.WebhookParams{
			Content:   services.AutoEmbed(m.Content),
			Username:  m.Author.Username,
			AvatarURL: m.Author.AvatarURL(""),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "❌",
							Style:    discordgo.SecondaryButton,
							CustomID: "delete_" + m.Author.ID + "_" + m.Content,
						},
					},
				},
			},
		})
		
		s.WebhookDelete(webhook.ID)
	}
}

func HandleDelete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if strings.HasPrefix(i.MessageComponentData().CustomID, "delete_") {
		parts := strings.Split(i.MessageComponentData().CustomID, "_")
		if len(parts) >= 3 && parts[1] == i.Member.User.ID {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredMessageUpdate,
			})
			s.ChannelMessageDelete(i.ChannelID, i.Message.ID)
		}
	}
}

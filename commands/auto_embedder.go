package commands

import (
	"discord_bot/services"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var domains = []string{"instagram.com", "pixiv.com", "x.com", "tiktok.com", "imgur.com", "twitter.com", "reddit.com", "facebook.com"}

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
		s.ChannelMessageSend(m.ChannelID, services.AutoEmbed(m.Content))
		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}
}

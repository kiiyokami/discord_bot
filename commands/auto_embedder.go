package commands

import (
	"discord_bot/services"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var domains = map[string]bool{
	"instagram.com": true, "pixiv.com": true, "x.com": true, "tiktok.com": true,
	"imgur.com": true, "twitter.com": true, "reddit.com": true,
}

func AutoEmbedder(s *discordgo.Session, m *discordgo.MessageCreate) {
	if u, err := url.ParseRequestURI(m.Content); err == nil {
		host := strings.TrimPrefix(u.Host, "www.")
		if domains[host] {
			embeddedURL := services.AutoEmbed(m.Content)
			s.ChannelMessageSend(m.ChannelID, embeddedURL)
			s.ChannelMessageDelete(m.ChannelID, m.ID)
		}
	}
}

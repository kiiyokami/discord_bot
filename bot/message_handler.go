package bot

import (
	"discord_bot/commands"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	commands.AutoEmbedder(s, m)
}

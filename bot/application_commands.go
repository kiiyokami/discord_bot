package bot

import (
	"discord_bot/commands"
	"log"

	"github.com/bwmarrin/discordgo"
)

var CommandsList = []*discordgo.ApplicationCommand{
	{
		Name:        "convert-time",
		Description: "Converts time from one timezone to another",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "hour",
				Description: "The hour to convert",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "minute",
				Description: "The minute to convert",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "country",
				Description: "The basis country to convert to",
				Required:    true,
			},
		},
	},
	{
		Name:        "fetch-image",
		Description: "Fetches a random image",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "image",
				Description: "The folder to search for",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Milkteaism",
						Value: "milktea",
					},
				},
			},
		},
	},
}

func RegisterCommands(s *discordgo.Session) {
	for _, cmd := range CommandsList {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", cmd.Name, err)
		}
	}
}

func CommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "convert-time":
		commands.ConvertTime(s, i)
	case "fetch-image":
		commands.FetchImage(s, i)
	}

}

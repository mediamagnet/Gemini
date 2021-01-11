package commands

import (
	"Gemini/lib"
	"github.com/bwmarrin/discordgo"
)

func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == lib.Prefix()+"help" {
		if m.Author.ID == "639949497467797524" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "]help")
		} else {
			_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title: "Gemini Help:",
				Author: &discordgo.MessageEmbedAuthor{
					URL:     "https://github.com/mediamagnet/gemini",
					Name:    "Gemini",
					IconURL: "https://cdn.discordapp.com/avatars/783065070682243082/418be8dcc03596073565a8706ed519ec.png?size=128"},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL:    "https://cdn.discordapp.com/avatars/783065070682243082/418be8dcc03596073565a8706ed519ec.png?size=128",
					Width:  128,
					Height: 128,
				},
				Color:       0x0047AB,
				Description: "Welcome to the Gemini, Here's some useful commands: \n",
				Fields: []*discordgo.MessageEmbedField{
					{Name: lib.Prefix()+"help", Value: "You're reading it."},
					{Name: lib.Prefix()+"role", Value: "Various commands for managing roles."},
					{Name: lib.Prefix()+"ping", Value: "Get ping time between Gemini and other parts of the bot."},
					{Name: lib.Prefix()+"info", Value: "Get info about this half of the bot."},
					{Name: lib.Prefix()+"register", Value: "(Pending), register your assigned roles with the bot"},
					{Name: lib.Prefix()+"login", Value: "(Pending), login to your roles for pings"},
					{Name: lib.Prefix()+"logout", Value: "(Pending), logout of your roles."},
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Gemini: the other half of the battle.",
					IconURL: "https://cdn.discordapp.com/avatars/783065070682243082/418be8dcc03596073565a8706ed519ec.png?size=16",
				},
			},
			)
		}
	}
}
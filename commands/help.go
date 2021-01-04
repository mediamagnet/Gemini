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
					{Name: ".command", Value: "Help text here"},
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
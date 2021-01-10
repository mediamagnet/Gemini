package commands

import (
	"Gemini/lib"
	"github.com/bwmarrin/discordgo"
	"runtime"
	"strconv"
)

func HelpInfo() CommandHelp {
	return CommandHelp{
		Name:    "Info",
		Usage:   lib.Prefix()+"info",
		Description: "Provides info about the bot",
		Admin:   false,
	}
}
func InfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	version := runtime.Version()
	if m.Content == lib.Prefix()+"info" {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "Info Panel",
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
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Gemini Go Version", Value: "v0.0.1"},
				{Name: "Go Version", Value: version, Inline: true},
				{Name: "Discord-Go Version", Value: discordgo.VERSION, Inline: true},
				{Name: "API version", Value: discordgo.APIVersion},
				{Name: "Shards", Value: strconv.Itoa(s.ShardCount), Inline: true},

			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "Gemini: the other half of the battle.",
				IconURL: "https://cdn.discordapp.com/avatars/783065070682243082/418be8dcc03596073565a8706ed519ec.png?size=16",
			},
		})
	}
}

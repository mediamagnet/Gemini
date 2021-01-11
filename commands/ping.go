package commands

import (
	"Gemini/config"
	"Gemini/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/go-ping/ping"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

var pingMax []string

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	var cfg config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Fatalf("unable to decode into struct, %v", err)
	}
	// msg := m.Content
	if strings.Contains(m.Content, lib.Prefix()+"ping") {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			log.Warnln(err)
		}

		msg1, _ := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Color:       0x00FF00,
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Discord", Value: "Pinging..."},
				{Name: "Mongo", Value: "Pinging..."},
				{Name: "Other Half", Value: "Pinging..."},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "Gemini: the other half of the battle.",
				IconURL: "https://cdn.discordapp.com/avatars/783065070682243082/418be8dcc03596073565a8706ed519ec.png?size=16",
			},
		})

		addrs := [2]string{"discord.com", "gemini-shard-00-02.hjehy.mongodb.net"}
		for _, s := range addrs {
			pinger, err := ping.NewPinger(s)
			pinger.SetPrivileged(true)
			if err != nil {
				panic(err)
			}
			pinger.Count = 3
			pinger.OnFinish = func(stats *ping.Statistics) {
				fmt.Println(stats)
				pingMax = append(pingMax, stats.AvgRtt.String())
			}
			_ = pinger.Run() // blocks until finished
		}
		_, _ = s.ChannelMessageEditEmbed(m.ChannelID, msg1.ID, &discordgo.MessageEmbed{
			Color:       0x0047AB,
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Discord", Value: pingMax[0]},
				{Name: "Mongo", Value: pingMax[1]},
				{Name: "Other Half", Value: "Watching the world go by"},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "Gemini: the other half of the battle.",
				IconURL: "https://cdn.discordapp.com/avatars/783065070682243082/418be8dcc03596073565a8706ed519ec.png?size=16",
			},
		})
	}
}

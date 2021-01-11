package commands

import (
	"Gemini/config"
	"Gemini/lib"
	"github.com/bwmarrin/discordgo"
	"github.com/go-ping/ping"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

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
	msg := m.Content
	if strings.HasPrefix(msg, lib.Prefix()+"ping") {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			log.Errorln(err)
		}
		disping, _ := ping.NewPinger("discord.com")
		monping, _ := ping.NewPinger("gemini-shard-00-02.hjehy.mongodb.net")
		gemping, _ := ping.NewPinger("geminibot.xyz")

/*		msg1, _ := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Discord", Value: "Pinging..."},
				{Name: "Mongo", Value: "Pinging..."},
				{Name: "Other Half", Value: "Pinging..."},
			},
		})*/

		disping.Count = 3
		monping.Count = 3
		gemping.Count = 3
		_ = disping.Run()
		_ = monping.Run()
		_ = gemping.Run()
		disStats := disping.Statistics()
		monStats := monping.Statistics()
		gemStats := gemping.Statistics()
		log.Infoln(disStats)
		log.Infoln(monStats)
		log.Infoln(gemStats)

/*		_, _ = s.ChannelMessageEditEmbed(m.ChannelID, msg1.ID, &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Discord", Value: disStats.MaxRtt.String()},
				{Name: "Mongo", Value: monStats.MaxRtt.String()},
				{Name: "Other Half", Value: gemStats.MaxRtt.String()},
			},
		})*/
	}
}

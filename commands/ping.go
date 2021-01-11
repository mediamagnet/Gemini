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
	// msg := m.Content
	if strings.Contains(m.Content, lib.Prefix()+"ping") {
		err := s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			log.Warnln(err)
		}
		disping, err := ping.NewPinger("www.discord.com")
		if err != nil {
			log.Errorf("b %v", err)
		}
		monping, err := ping.NewPinger("gemini-shard-00-02.hjehy.mongodb.net")
		if err != nil {
			log.Errorf("c %v", err)
		}
		gemping, err := ping.NewPinger("geminibot.xyz")
		if err != nil {
			log.Errorf("a %v", err)
		}

/*		msg1, _ := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Discord", Value: "Pinging..."},
				{Name: "Mongo", Value: "Pinging..."},
				{Name: "Other Half", Value: "Pinging..."},
			},
		})*/

		disping.SetPrivileged(true)
		monping.SetPrivileged(true)
		gemping.SetPrivileged(true)

		disping.Count = 3
		monping.Count = 3
		gemping.Count = 3

		err = disping.Run()
		if err != nil {
			log.Errorf("d %v", err)
		}
		err = monping.Run()
		if err != nil {
			log.Errorf("e %v", err)
		}
		err = gemping.Run()
		if err != nil {
			log.Errorf("f %v", err)
		}
		// disStats := disping.Statistics()
		// monStats := monping.Statistics()
		// gemStats := gemping.Statistics()
		log.Infof("a, %v", disping.Statistics())
		log.Infof("b, %v", monping.Statistics())
		log.Infof("c, %v", gemping.Statistics())

/*		_, _ = s.ChannelMessageEditEmbed(m.ChannelID, msg1.ID, &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{
				{Name: "Discord", Value: disStats.MaxRtt.String()},
				{Name: "Mongo", Value: monStats.MaxRtt.String()},
				{Name: "Other Half", Value: gemStats.MaxRtt.String()},
			},
		})*/
	}
}

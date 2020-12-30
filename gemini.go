package main

import (
	"Gemini-Go/config"
	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	var cfg config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Fatal("unable to decode into struct, %v", err)
	}

	client := atlas.New(&atlas.Options{
		DisgordOptions: disgord.Config{
			BotToken: cfg.Bot.Token,
			Logger:   log,
		},
		OwnerID: cfg.Owner.ID,
	})

	client.Use(atlas.DefaultLogger())
	client.GetPrefix = func(m *disgord.Message) string {
		return cfg.Bot.Prefix
	}

	if err := client.Init(); err != nil {
		panic(err)
	}

}

func init() {
	atlas.Use(commands.InitRole().Register())
	atlas.Use(commands.InitHelp().Register())

}
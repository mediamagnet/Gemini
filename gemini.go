package main

import (
	"Gemini/commands"
	"Gemini/config"
	"Gemini/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

//Prefix is the bots prefix for commands
var Prefix string

func main() {

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
	Prefix = lib.Prefix()

	dg, err := discordgo.New("Bot " + cfg.Bot.Token)
	if err != nil {
		log.Fatalf("Error creating Discord session, %v", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(connect)

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection, %v", err)
		return
	}
	log.Infoln("Bot is now running. Press CTRL-C to exit.")


	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()

}
func connect(s *discordgo.Session, c *discordgo.Connect) {
	log.Println(c)
	var guildName = make([]string, 1)
	for _, v := range s.State.Guilds {
		guildName = append(guildName, v.Name)
	}
	for {
		err := s.UpdateListeningStatus("cosmic background radiation")
		time.Sleep(15 * time.Minute)
		err = s.UpdateStatus(0, "Gemini v0.0.1")
		time.Sleep(15 * time.Minute)
		err = s.UpdateListeningStatus(".help")
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(15 * time.Minute)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Test to make sure bot isn't talking to self.
	if m.Author.ID == s.State.User.ID {
		return
	}
	switch {
	case lib.ComesFromDM(s, m) == false:
		commands.HelpCommand(s, m)
		commands.CleanupCommand(s, m)
		commands.RoleCommand(s, m)
	case lib.ComesFromDM(s, m) == true && m.Author.ID == "108344508940316672":
		log.Printf("It works Prefix is: %s", Prefix)
	default:
		log.Println("It's a DM")
	}
}
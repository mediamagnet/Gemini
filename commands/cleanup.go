package commands

import (
	"Gemini/lib"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}
// var wg sync.WaitGroup

//CleanupCommand provides channel purge functionality
func CleanupCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, lib.Prefix()+"cleanup") {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
		delCount1 := strings.TrimPrefix(m.Content, lib.Prefix()+"cleanup ")
		delCount, _ := strconv.Atoi(delCount1)
		log.Println(delCount)
		if lib.MemberHasPermission(s, m.GuildID, m.Author.ID, discordgo.PermissionManageMessages|discordgo.PermissionAdministrator) {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Deleting messages in <#"+m.ChannelID+">")
			if delCount == 0 {
				time.Sleep(3 * time.Second)
				// wg.Add(100)
				log.Println("Cleanup Requested")
				for i := 0; i < 1000; i++ {
					messages, _ := s.ChannelMessages(m.ChannelID, 1, "", "", "")
					if len(messages) == 0 {
						break
					}
					// fmt.Println(messages[0].ID)
					println(i)
					time.Sleep(250 * time.Millisecond)
					_ = s.ChannelMessageDelete(m.ChannelID, messages[0].ID)
					log.Println("done cleaning", i)
				}
			} else {
				time.Sleep(3 * time.Second)
				// wg.Add(100)
				log.Println("Cleanup Requested")
				for i := 0; i < delCount+1; i++ {
					messages, _ := s.ChannelMessages(m.ChannelID, 1, "", "", "")
					if len(messages) == 0 {
						break
					}
					log.Println(messages[0].ID)
					time.Sleep(250 * time.Millisecond)
					_ = s.ChannelMessageDelete(m.ChannelID, messages[0].ID)
					log.Println("done cleaning", i)
				}
			}
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Sorry <@"+m.Author.ID+"> You need Manage Message permissions to run .cleanup")
		}
	}
}

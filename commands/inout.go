package commands

import (
	"Gemini/lib"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func InOutCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Content
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	var roleStr []string
	if err != nil {
		if member, err = s.GuildMember(m.GuildID, m.Author.ID); err != nil {
			log.Errorln(err)
		}
	}
	switch {
		case strings.Contains(msg, lib.Prefix()+"register"):
			_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
			for _, roleID := range member.Roles {
				role, err := s.State.Role(m.GuildID, roleID)
				if err != nil {
					log.Errorln(err)
				}
				roleStr = append(roleStr, string(role.ID))
			}
			userInput := lib.User{
				GuildID: m.GuildID,
				UserID: m.Author.ID,
				RoleIDs: roleStr,
			}
			lib.MonUser("gemini", "Users", userInput)
	}
}

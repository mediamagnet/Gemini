package commands

import (
	"Gemini/lib"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)


func RoleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Content
	if strings.HasPrefix(m.Content, lib.Prefix()+"role") {
		if lib.MemberHasPermission(s, m.GuildID, m.Author.ID, discordgo.PermissionManageRoles|discordgo.PermissionAdministrator) {
			switch {
			case strings.Contains(msg, lib.Prefix()+"role new"):
				msg = strings.TrimPrefix(msg, lib.Prefix()+"role new ")
				msg1 := strings.Split(msg, ", ")
				message, _ := s.ChannelMessageSend(m.ChannelID, "Creating...")
				roleInsert := lib.Role{
					GuildID:   m.GuildID,
					ChannelID: strings.TrimPrefix(strings.TrimSuffix(msg1[1], ">"), "<#"),
					RoleID:    strings.TrimPrefix(strings.TrimSuffix(msg1[0], ">"), "<@&"),
					Phrase:    msg1[2],
				}
				lib.MonRole("gemini", "roles", roleInsert)
				fmt.Println(message.ID)
				_, err := s.ChannelMessageEdit(m.ChannelID, message.ID, fmt.Sprintf("Done adding role, %v", msg1[0]))
				if err != nil {
					log.Warnln(err)
				}
			case strings.Contains(msg, lib.Prefix()+"role react"):
				log.Println("soon")
				msg = strings.TrimPrefix(msg, lib.Prefix()+"role react ")
				msg1 := strings.Split(msg, ", ")
				lib.MonUpdateRole(lib.GetClient(), bson.M{"Reaction": msg1[2], "Message": msg1[1]}, bson.M{"RoleID": strings.TrimPrefix(strings.TrimSuffix(msg1[0], ">"), "<@&")})
				err := s.MessageReactionAdd(m.ChannelID, msg1[2],msg1[1])
				if err != nil {
					log.Warnln(err)
				}
			case strings.Contains(msg, lib.Prefix()+"role get"):
				msg := strings.TrimPrefix(msg, lib.Prefix()+"role get ")
				msg1 := strings.Split(msg, ", ")
				fmt.Println(s.MessageReactions(m.ChannelID, msg1[0], msg1[1], 100, m.ID, msg1[2]))
			default:
				_, err := s.ChannelMessageSend(m.ChannelID, "To add a role assignment, please respond with the following: "+lib.Prefix()+"role create <channel id> <role id>")
				if err != nil {
					log.Warnln(err)
				}
			}
		}
	}
}

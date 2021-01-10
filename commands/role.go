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
		msg3 := strings.TrimPrefix(msg, lib.Prefix()+"role ")
		var phrase string
		var roleID string
		phrases := lib.MonReturnAllRecords(lib.GetClient(), bson.M{}, "gemini", "roles")
		phraseLookUp := lib.MonReturnOneRecord(lib.GetClient(), bson.M{"GuildID": m.GuildID}, "gemini", "roles")
		log.Infof("%v, %v", phrases, phraseLookUp)
		if m.GuildID == phraseLookUp.GuildID {
			for _, v := range phrases {
				if v.Phrase == msg3 {
					log.Infof(v.Phrase)
					phrase = v.Phrase
					roleID = v.RoleID
				}
			}
		}
		if msg3 == phrase {
			_, err := s.ChannelMessageSend(m.ChannelID, "Adding role <@&"+roleID+"> to "+m.Author.Mention())
			if err != nil {
				log.Warnln(err)
			}
			err = s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, roleID)
			if err != nil {
				log.Warnln(err)
			}
		} else if lib.MemberHasPermission(s, m.GuildID, m.Author.ID, discordgo.PermissionManageRoles|discordgo.PermissionAdministrator) {
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
				lib.MonRecord("gemini", "roles", roleInsert)
				fmt.Println(message.ID)
				_, err := s.ChannelMessageEdit(m.ChannelID, message.ID, fmt.Sprintf("Done adding role, %v", msg1[0]))
				if err != nil {
					log.Warnln(err)
				}
			case strings.Contains(m.Content, lib.Prefix()+"role drop"):
				msg := strings.TrimPrefix(msg, lib.Prefix()+"role drop ")
				msg1 := strings.TrimPrefix(strings.TrimSuffix(msg, ">"), "<@&")
				message, err := s.ChannelMessageSend(m.ChannelID, "Dropping role from monitoring...")
				if err != nil {
					log.Warnln(err)
				}
				lib.MonDeleteRecord(lib.GetClient(), bson.M{"RoleID": msg1}, "gemini", "roles")
				_, err = s.ChannelMessageEdit(m.ChannelID, message.ID, fmt.Sprintf("No longer watching for role: %v", msg))
				if err != nil {
					log.Warnln(err)
				}
			default:
				_, err := s.ChannelMessageSend(m.ChannelID, "To add a role assignment, please respond with the following: `"+lib.Prefix()+"role new <channel id>, <role id>, <phrase>`")
				if err != nil {
					log.Warnln(err)
				}
			}
		}
	}
}

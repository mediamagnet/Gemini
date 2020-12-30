package commands

import (
	"Bolts/lib"
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/andersfylling/disgord"

	"strings"

	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// Role Struct
type Role struct{ Command }

// InitRole sets up the role command
func InitRole() Role {
	return Role{Init(&CommandItem{
		Name:        "role",
		Description: "Give User a role when a phrase is said",
		Usage:       "To create a role '" + os.Getenv("PREFIX") + " new <listenchannel> <role>, <phrase>' (Requires Manage Roles Permissions) To assign a role: '" + os.Getenv("PREFIX") + " role <phrase>'",
		Parameters: []Parameter{
			{
				Name:        "listenchannel",
				Description: "Channel to listen for phrase in.",
				Required:    true,
			},
			{
				Name:        "Role",
				Description: "Role to give",
				Required:    true,
			},
			{
				Name:        "Ignore",
				Description: "Role to Ignore",
				Required:    true,
			},
			{
				Name:        "Phrase",
				Description: "Phrase to listen for.",
				Required:    true,
			},
		},
		Admin: false,
	})}
}

// Register command
func (c Role) Register() *atlas.Command {
	// var guildFound string
	// var role disgord.Snowflake
	// var channel disgord.Snowflake
	var phrase string
	var roleID = ""
	var ignoreID = ""
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s]+")
	if err != nil {
		logrus.Fatal(err)
	}

	c.CommandInterface.Run = func(ctx atlas.Context) {

		guildString := ctx.Message.GuildID.String()

		phraseLookUp := lib.MonReturnOneListen(lib.GetClient(), bson.M{"GuildID": guildString})
		var chanIDClean string
		var roleIDClean string
		var ignoreIDClean string

		msg := strings.ToLower(reg.ReplaceAllString(strings.TrimPrefix(ctx.Message.Content, os.Getenv("PREFIX")+"role "), ""))
		p, err := disgord.Session.GetMemberPermissions(ctx.Atlas.Disgord, context.Background(), ctx.Message.GuildID, ctx.Message.Author.ID)
		fmt.Printf("MSG: %s \n", msg)

		phrases := lib.MonReturnAllListen(lib.GetClient(), bson.M{})
		if ctx.Message.GuildID.String() == phraseLookUp.GuildID {
			for _, v := range phrases {
				if v.Phrase == msg {
					phrase = v.Phrase
					roleID = v.RoleID
					ignoreID = v.IgnoreID
				}
			}
		}

		fmt.Printf("Phrase: %s \n", phrase)

		fmt.Println(phrase)

		phraseClean := strings.ToLower(reg.ReplaceAllString(phrase, ""))

		if msg == phraseClean {
			userroles, _ := atlas.Disgord.GetMember(ctx.Atlas.Disgord, context.TODO(), ctx.Message.GuildID, ctx.Message.Author.ID)
			rolestring := fmt.Sprintf("%v", userroles.Roles)
			_ = atlas.Disgord.DeleteMessage(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, ctx.Message.ID)
			if strings.Contains(rolestring, ignoreID) {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "Sorry I am not able to give you that role at current. Someone will be along shortly to help out.")
			} else {
				fmt.Printf("Role: %s \n", roleID)

				roleStr := lib.StrToSnowflake(roleID)
				fmt.Printf("Snowflake: %v, Converted: %v \n", roleStr, lib.SnowflakeToUInt64(roleStr))

				err = atlas.Disgord.AddGuildMemberRole(ctx.Atlas.Disgord, context.Background(), ctx.Message.GuildID, ctx.Message.Author.ID, roleStr)
				if err != nil {
					logrus.Warn(err)
				}
				msgstr := fmt.Sprintf("Acknowledged, %v", ctx.Message.Author.Username)
				_, _ = ctx.Message.Reply(ctx.Context, ctx.Atlas, msgstr)
			}
		} else {
			if len(ctx.Args) > 0 {
				phrase := strings.TrimPrefix(ctx.Message.Content, os.Getenv("PREFIX")+"role")
				phrase1 := strings.Split(phrase, ", ")
				fmt.Println(ctx.Args[0])
				if ctx.Args[0] == "new" {
					if p&disgord.PermissionManageRoles == 0 {
						ctx.Message.Reply(ctx.Context, ctx.Atlas, "Sorry you don't have the required permissions to create a new role.")
					} else {
						replyPhrase := fmt.Sprintf("Watching for %v%v", os.Getenv("PREFIX"), phrase1[len(phrase1)-1])
						ctx.Message.Reply(ctx.Context, ctx.Atlas, replyPhrase)
						fmt.Println(phrase)
						fmt.Printf("Guild: %v, Channel: %v, Add: %v, Ignore: %v, Phrase: %v \n", ctx.Message.GuildID, ctx.Args[0], ctx.Args[1], ctx.Args[2], phrase1[len(phrase1)-1])

						// Allow for both ID and Channel name
						if strings.HasPrefix(ctx.Args[1], "<#") {
							chanIDClean = strings.TrimPrefix(strings.TrimSuffix(ctx.Args[1], ">"), "<#")
						} else {
							chanIDClean = strings.TrimSuffix(ctx.Args[1], ">")
						}

						// Allow for both ID and Role name
						if strings.HasPrefix(ctx.Args[2], "<@&") {
							roleIDClean = strings.TrimPrefix(strings.TrimSuffix(ctx.Args[2], ">"), "<@&")
						} else {
							roleIDClean = strings.TrimSuffix(ctx.Args[2], ">")
						}

						// Allow Ignored role id and name
						if strings.HasPrefix(ctx.Args[2], "<@&") {
							ignoreIDClean = strings.TrimPrefix(strings.TrimSuffix(ctx.Args[3], ">,"), "<@&")
						} else {
							ignoreIDClean = strings.TrimSuffix(ctx.Args[3], ">,")
						}

						listenInsert := lib.RoleMeListen{
							GuildID:   ctx.Message.GuildID.String(),
							ChannelID: chanIDClean,
							RoleID:    roleIDClean,
							IgnoreID:  ignoreIDClean,
							Phrase:    strings.ToLower(reg.ReplaceAllString(phrase1[len(phrase1)-1], "")),
						}
						lib.MonListen("bolts", "listens", listenInsert)
					}
				}
			} else {
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "Sorry, it seems you're missing a role phrase or the rest of the command to add a role. If you need help do ]h role.")
				// member, err := disgord.Session.GetMember(ctx.Atlas.Disgord, context.Background(),ctx.Message.GuildID, ctx.Message.Author.ID)
				// if err != nil {
				// 	logrus.Fatal(err)
				// }

				fmt.Println(disgord.Session.GetMember(ctx.Atlas.Disgord, context.TODO(), ctx.Message.GuildID, ctx.Message.Author.ID))

				fmt.Printf("context: %v \n", ctx.Atlas.Disgord)
				if err != nil {
					logrus.Fatal(err)
				}

				fmt.Printf("canDo: %v \n", p)
				fmt.Printf("Admin: %v \n", p&disgord.PermissionManageRoles)

			}
		}
	}
	return c.CommandInterface
}

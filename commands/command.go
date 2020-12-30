package commands

import (
	"Bolts/lib"
	"os"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
)

// Command is the base command object for all commands.
type Command struct {
	CommandInterface *atlas.Command
}

// CommandItem is the base command item object for the help command.
type CommandItem struct {
	Name        string
	Description string
	Aliases     []string
	Usage       string
	Parameters  []Parameter
	Admin       bool
}

// Parameter is the base parameter object for the help command.
type Parameter struct {
	Name        string // parameter name
	Value       string // value representation
	Description string // parameter description
	Required    bool   // is parameter required?
}

// Init initializes atlas commands
func Init(t *CommandItem) Command {
	cmd := atlas.NewCommand(t.Name).SetDescription(t.Description)

	if t.Aliases != nil {
		cmd.SetAliases(t.Aliases...)
	}

	commands = append(commands, *t)

	return Command{cmd}
}

// getBot returns the bot object.
func (c Command) getBot(ctx atlas.Context) *disgord.Member {
	id, _ := strconv.Atoi(os.Getenv("BOT_ID"))
	bot, _ := ctx.Atlas.GetMember(ctx.Context, ctx.Message.GuildID, disgord.NewSnowflake(uint64(id)))

	return bot
}

// getBotUser returns the bot User object.
func (c Command) getBotUser(ctx atlas.Context) *disgord.User {
	return c.getBot(ctx).User
}

func (c Command) embedFooter(ctx atlas.Context) (f *disgord.EmbedFooter, t disgord.Time) {
	f, t = lib.AddEmbedFooter(ctx.Message)

	return
}

package commands

import (
	"fmt"
	"os"
	"github.com/pazuzu156/atlas"
	"strings"
	"time"
)

// Ping is a simple testing command.
type Ping struct{ Command }

// InitPing initializes the ping command.
func InitPing() Ping {
	return Ping{Init(&CommandItem{
		Name:        "ping",
		Description: "Ping/Pong",
		Usage:       "ping Pong!",
		Parameters: []Parameter{
			{
				Name:        "string",
				Description: "A string to send back to yourself",
				Required:    false,
			},
		},
	})}
}

// Register registers and runs the ping command.
func (c Ping) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		cmdTime := time.Now()
		if ctx.Message.Author.ID.String() != os.Getenv("BOT_ID") {
			if len(ctx.Args) > 0 {
				msg := strings.TrimPrefix(ctx.Message.Content, "]ping ")
				fmt.Println(ctx.Message.Content)
				ctx.Message.Reply(ctx.Context, ctx.Atlas, msg)
			} else {
				time2 := time.Now()
				difference := time2.Sub(cmdTime)
				ctx.Message.Reply(ctx.Context, ctx.Atlas, "Pong "+difference.String())
			}
		}
	}

	return c.CommandInterface
}

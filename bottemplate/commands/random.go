package commands

import (
	"Melvin/bottemplate"
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"math/rand"
)

var random = discord.SlashCommandCreate{
	Name:        "random",
	Description: "generate a random number from 0 to x",
}

func RandomHandler(b *bottemplate.Bot) handler.CommandHandler {
	// data := event.SlashCommandInteractionData()
	roll := generateRandom(0, 100)
	return func(e *handler.CommandEvent) error {
		return e.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("your roll: %d", roll),
		})
	}
}

func generateRandom(min, max int) int {
	return rand.Intn(max-min+1) + min
}

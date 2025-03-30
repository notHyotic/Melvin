package commands

import (
	"Melvin/bottemplate"
	"fmt"
	"strconv"

	"math/rand"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var roll = discord.SlashCommandCreate{
	Name:        "roll",
	Description: "generate a random number from 0 to x (default is 100)",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:         "max",
			Description:  "max number result can be",
			Required:     false,
			Autocomplete: false,
		},
	},
}

func RollHandler(b *bottemplate.Bot) handler.CommandHandler {
	// data := event.SlashCommandInteractionData()
	return func(e *handler.CommandEvent) error {
		// Default max value
		maxVal := 100

		// Get interaction data
		data := e.SlashCommandInteractionData()

		// Check if "max" was provided
		if option, ok := data.OptString("max"); ok {
			// Convert from string to int
			if parsedMax, err := strconv.Atoi(option); err == nil {
				maxVal = parsedMax
			} else {
				return e.CreateMessage(discord.MessageCreate{
					Content: "Invalid number format for max. Please enter a valid integer.",
				})
			}
		}
	
		roll := generateRandom(0, maxVal)

		return e.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("your roll: %d", roll),
		})
	}
}

func generateRandom(min, max int) int {
	return rand.Intn(max-min+1) + min
}

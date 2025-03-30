package commands

import (
	"Melvin/bottemplate"
	"fmt"
	"math/big"
	"strconv"

	"crypto/rand"

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
	
		roll, err := generateRandom(0, maxVal)
		if err != nil {
			fmt.Println("Error:", err)
			return e.CreateMessage(discord.MessageCreate{
				Content: err.Error(),
			})
		}
		return e.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("your roll: %d", roll),
		})
	}
}

// generateRandom securely generates a random integer between min and max (inclusive)
func generateRandom(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min must be less than or equal to max")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + min, nil
}

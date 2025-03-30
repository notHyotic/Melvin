package handlers

import (
	"Melvin/bottemplate"
	"fmt"
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
)

func UserJoinVCHandler(b *bottemplate.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.GuildVoiceJoin) {
		accId := 840529939139919874
		channel, err := bot.Client.
			Rest(b.Client).
			CreateDMChannel(snowflake.ID(accId))
		if err != nil {
			slog.Error("Error occured while trying to create DM channel", err)
			return
		}

		userID := &e.VoiceState.UserID
		guildID := &e.VoiceState.GuildID
		channelID := e.VoiceState.ChannelID

		msg, err := createNotificationMessage(b, *userID, *guildID, *channelID)
		if err != nil {
			fmt.Errorf("Error occurred while trying to create notification message %w", err)
			return
		}


		if rune(*userID) == rune(accId) {
			return
		}
		bot.Client.Rest(b.Client).CreateMessage(
			channel.ID(),
			discord.
				NewMessageCreateBuilder().
				SetContent(msg).
				Build())
	})
}

func createNotificationMessage(b *bottemplate.Bot, userId, guildId, channelId snowflake.ID) (string, error) {
	user, err := b.Client.Rest().GetMember(guildId, userId)
	if err != nil {
		slog.Error("Error occured while trying to get user info")
		return "", err
	}

	channel, err := b.Client.Rest().GetChannel(channelId)
	if err != nil {
		slog.Error("Error occured while trying to get channel info")
		return "", err
	}

	message := user.User.Username + " has joined voice channel " + channel.Name()

	return message, nil
}

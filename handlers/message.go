package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/john-d-pelingo/godbot/commands"
	"github.com/john-d-pelingo/godbot/helpers"
)

// HandleMessage handles the incoming message and lets the handlers manage it accordingly.
func HandleMessage(prefix string) func(discord *discordgo.Session, message *discordgo.MessageCreate) {
	return func(discord *discordgo.Session, message *discordgo.MessageCreate) {
		user := message.Author

		if user.ID == discord.State.User.ID || user.Bot {
			return
		}

		content := message.Content

		if content == fmt.Sprintf("%sping", prefix) {
			commands.Ping(discord, message)
		}

		helpers.PrettyPrint(&message.Message)
	}
}

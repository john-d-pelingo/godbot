package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/john-d-pelingo/godbot/commands"
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

		if content == fmt.Sprintf("%shelp", prefix) {
			commands.Help(discord, message, commands.HelpOptions{
				Prefix: prefix,
			})
		}

		prettyPrint(&message.Message)
	}
}

// prettyPrint prints the contents of data with spaces for better distinction.
func prettyPrint(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix string
	botID         string
)

func main() {
	config := initConfig("./config.json")

	dgo, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	errCheck("Failed to create discord session.", err)

	bot, err := dgo.User("@me")
	errCheck("Failed to access account.", err)

	botID = bot.ID
	dgo.AddHandler(handleCmd)

	err = dgo.Open()
	errCheck("Unable to establish connection.", err)

	defer dgo.Close()

	run()
}

func handleCmd(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author

	if user.ID == botID || user.Bot {
		return
	}

	content := message.Content

	if content == "!test" {
		discord.ChannelMessageSend(message.ChannelID, "Testing something...")
	}

	prettyPrint(&message.Message)
}

package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	conf "github.com/john-d-pelingo/godbot/config"
	"github.com/john-d-pelingo/godbot/helpers"
)

var (
	commandPrefix string
	botID         string
)

func main() {
	config := conf.InitConfig("./config.json")

	dgo, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	helpers.ErrCheck("Failed to create discord session.", err)

	bot, err := dgo.User("@me")
	helpers.ErrCheck("Failed to access account.", err)

	botID = bot.ID
	dgo.AddHandler(handleCmd)

	err = dgo.Open()
	helpers.ErrCheck("Unable to establish connection.", err)

	defer dgo.Close()

	helpers.Run()
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

	helpers.PrettyPrint(&message.Message)
}

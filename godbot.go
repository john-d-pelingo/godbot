package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	conf "github.com/john-d-pelingo/godbot/config"
	"github.com/john-d-pelingo/godbot/helpers"
)

func main() {
	config := conf.InitConfig("./config.json")

	dgo, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	helpers.ErrCheck("Failed to create discord session.", err)

	_, err = dgo.User("@me")
	helpers.ErrCheck("Failed to access account.", err)

	dgo.AddHandler(handleCmd)

	err = dgo.Open()
	helpers.ErrCheck("Unable to establish connection.", err)

	defer dgo.Close()

	helpers.Run()
}

func handleCmd(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author

	if user.ID == discord.State.User.ID || user.Bot {
		return
	}

	content := message.Content

	if content == "!test" {
		_, err := discord.ChannelMessageSend(message.ChannelID, "Testing something...")

		if err != nil {
			fmt.Println(err)
		}
	}

	helpers.PrettyPrint(&message.Message)
}

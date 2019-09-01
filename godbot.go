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
	discord, err := discordgo.New("Bot <BOT_TOKEN>")
	errCheck("failed to create discord session", err)

	bot, err := discord.User("@me")
	errCheck("failed to access account", err)

	botID = bot.ID
	discord.AddHandler(handleCmd)

	err = discord.Open()
	errCheck("unable to establish connection", err)

	defer discord.Close()

	<-make(chan struct{})
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

	fmt.Printf("Message: %+v\n", message.Message)
}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

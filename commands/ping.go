package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/john-d-pelingo/godbot/helpers"
)

// Ping is in charge of the ping command and sends a corresponding message back.
func Ping(discord *discordgo.Session, message *discordgo.MessageCreate) {
	timestamp, err := message.Timestamp.Parse()

	if err != nil {
		fmt.Println(err)
		return
	}

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{Name: "GodBot", URL: "https://github.com/john-d-pelingo/godbot"},
		Color:       helpers.RandomRGBColorNumber(),
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       "Pong!",
		Description: fmt.Sprintf("Your ping: %v", time.Since(timestamp)),
	}

	_, err = discord.ChannelMessageSendEmbed(message.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
	}
}

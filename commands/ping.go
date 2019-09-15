package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Ping is in charge of the ping command and sends a corresponding message back.
func Ping(discord *discordgo.Session, message *discordgo.MessageCreate) {
	timestamp, err := message.Timestamp.Parse()

	if err != nil {
		fmt.Println(err)
		return
	}

	ping := time.Since(timestamp)
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{Name: "GodBot", URL: "https://github.com/john-d-pelingo/godbot"},
		Color:       16098851, // TODO: generate random colors
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       "Pong!",
		Description: fmt.Sprintf("Your ping: %v", ping),
	}

	_, err = discord.ChannelMessageSendEmbed(message.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
	}
}

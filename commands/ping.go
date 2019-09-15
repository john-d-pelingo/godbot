package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Ping is in charge of  the ping command and sends a corresponding message back.
func Ping(discord *discordgo.Session, message *discordgo.MessageCreate) {
	timestamp, err := message.Timestamp.Parse()

	if err != nil {
		fmt.Println(err)
		return
	}

	ping := time.Since(timestamp)
	botInfo := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{Name: "GodBot", URL: "https://github.com/john-d-pelingo/godbot"},
		Color:       16098851,
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       "Pong!",
		Description: fmt.Sprintf("Your ping: %v", ping),
	}

	_, err = discord.ChannelMessageSendEmbed(message.ChannelID, botInfo)

	if err != nil {
		fmt.Println(err)
	}
}

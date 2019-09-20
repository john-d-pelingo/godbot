package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/john-d-pelingo/sinartisi/color"
)

// HelpOptions represents the specific options passed to the Help function.
type HelpOptions struct {
	Prefix string
}

// Help is in charge of the help command and sends a list of available commands back.
func Help(discord *discordgo.Session, message *discordgo.MessageCreate, options HelpOptions) {
	list := fmt.Sprintf(`Here is a list of commands below:
		%shelp - this command
		%sping - contact this bot
	`, options.Prefix, options.Prefix)

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{Name: "GodBot", URL: "https://github.com/john-d-pelingo/godbot"},
		Color:       color.RandomRGBColorNumber(),
		Timestamp:   time.Now().Format(time.RFC3339),
		Title:       "Help",
		Description: list,
	}

	_, err := discord.ChannelMessageSendEmbed(message.ChannelID, embed)
	if err != nil {
		fmt.Println(err)
	}
}

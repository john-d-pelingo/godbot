package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	conf "github.com/john-d-pelingo/godbot/config"
	"github.com/john-d-pelingo/godbot/handlers"
	"github.com/john-d-pelingo/godbot/helpers"
)

func main() {
	config := conf.InitConfig("./config.json")

	dgo, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	helpers.ErrCheck("Failed to create discord session.", err)

	_, err = dgo.User("@me")
	helpers.ErrCheck("Failed to access account.", err)

	go dgo.AddHandler(handlers.HandleMessage(config.Prefix))

	err = dgo.Open()
	helpers.ErrCheck("Unable to establish connection.", err)

	defer dgo.Close()

	helpers.Run()
}

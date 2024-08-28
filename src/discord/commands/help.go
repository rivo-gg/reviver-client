package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

const (
	help  = "help"
)

var commandHelpMap = map[string]string{
	help:  "https://reviverbot.com/commands",
}

var helpCommand = &discordgo.ApplicationCommand{
	Name:        "help",
	Description: "Displays command help information.",
}

func handleHelpCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	logrus.Info("Handling help command.")

	options := i.ApplicationCommandData().Options
	var doc string

	if len(options) == 0 {
		doc = fmt.Sprintf("Here's the documentation for Reviver: %s", commandHelpMap[help])
	} else {
		command := i.ApplicationCommandData().Options[0].StringValue()
		docLink := commandHelpMap[command]
		doc = fmt.Sprintf("Here's the documentation for the %s command: %s", command, docLink)
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: doc,
			Flags:   Ephemeral,
		},
	})
	if err != nil {
		sendError(help, err)
	}
}

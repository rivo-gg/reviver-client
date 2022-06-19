package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	discordgo_scm "github.com/ethanent/discordgo-scm"
	"github.com/servusdei2018/shards"
)

var SCM = discordgo_scm.NewSCM()

func RegisterCommands(s *shards.Manager, guild string) error {
	SCM.AddFeatures([]*discordgo_scm.Feature{
		{
			Type:               discordgo.InteractionApplicationCommand,
			Handler:            handleHelpCommand,
			ApplicationCommand: helpCommand,
		},
		{
			Type:               discordgo.InteractionApplicationCommand,
			Handler:            handleTopicCommand,
			ApplicationCommand: topicCommand,
		},
		{
			Type:               discordgo.InteractionApplicationCommand,
			Handler:            handleFactCommand,
			ApplicationCommand: factCommand,
		},
	})

	var session *discordgo.Session

	if guild != "" {
		guildId, err := strconv.ParseInt(guild, 10, 64)
		if err != nil {
			return err
		}

		session = s.SessionForGuild(guildId)
	} else {
		session = s.SessionForGuild(0)
	}

	return SCM.CreateCommands(session, guild)
}

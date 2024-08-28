package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/rivo-gg/reviver-go/src/discord/commands"
	"github.com/rivo-gg/reviver-go/src/discord/impl"
	"github.com/servusdei2018/shards"
	"github.com/sirupsen/logrus"
)

func onConnect(s *discordgo.Session, evt *discordgo.Connect) {
	logrus.Info("Connected on shard ", s.ShardID)
}

func StartDiscordClient() (error, *shards.Manager, chan commands.CommandError) {
	token := os.Getenv("TOKEN")
	guild := os.Getenv("GUILD")

	commands.Errors = make(chan commands.CommandError, 16)
	commands.Topics = impl.TopicManager{}

	commands.Topics.Load()

	sm, err := shards.New("Bot " + token)
	if err != nil {
		return err, nil, nil
	}

	sm.AddHandler(commands.SCM.HandleInteraction)
	sm.AddHandler(onConnect)

	err = sm.Start()
	if err != nil {
		return err, nil, nil
	}

	err = commands.RegisterCommands(sm, guild)
	if err != nil {
		return err, nil, nil
	}

	return nil, sm, commands.Errors
}

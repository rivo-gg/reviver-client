package commands

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/rivo-gg/reviver-go/src/database"
	"github.com/sirupsen/logrus"
)

var topicCommand = &discordgo.ApplicationCommand{
	Name:        "topic",
	Description: "Get a random topic to make chat more active!",
}

var factCommand = &discordgo.ApplicationCommand{
	Name:        "fact",
	Description: "Get a random fact to make chat more active!",
}

func createTopicEmbed(category database.Category, topic string, id int64, interaction *discordgo.Interaction) *discordgo.MessageEmbed {

	upperCategory := strings.ToUpper(string(category))

	embed := discordgo.MessageEmbed{
		Description: fmt.Sprintf("**%s**", topic),
		Color:       0x87CEEB,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Requested by %s | %s ID: ", interaction.Member.User.Username, upperCategory) + strconv.FormatInt(id, 10),
			IconURL: fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.webp", interaction.Member.User.ID, interaction.Member.User.Avatar),
		},
	}

	return &embed
}

func handleTopic(s *discordgo.Session, i *discordgo.InteractionCreate, category database.Category) {
	guildId, err := strconv.ParseInt(i.GuildID, 10, 64)
	if err != nil {
		sendError(fmt.Sprintf("cmd_%s", category), err)
		return
	}

	topicValue, err := Topics.GetRandomTopic(guildId, category)
	if err != nil {
		sendError(fmt.Sprintf("cmd_%s", category), err)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				createTopicEmbed(category, topicValue.Topic, topicValue.ID, i.Interaction),
			},
		},
	})
	if err != nil {
		sendError(fmt.Sprintf("cmd_%s", category), err)
	}
}

func handleTopicCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	logrus.Info("Handling topic command.")

	handleTopic(s, i, database.CategoryTopic)
}

func handleFactCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	logrus.Info("Handling fact command.")

	handleTopic(s, i, database.CategoryFact)
}

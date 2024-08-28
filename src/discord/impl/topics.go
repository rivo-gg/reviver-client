package impl

import (
	"math/rand"
	"time"

	"github.com/rivo-gg/reviver-go/src/database"
)

type Topic struct {
	ID    int64
	Topic string
}

type TopicManager struct {
	Topics       map[database.Category][]database.GlobalTopic
	ServerTopics map[int64]map[database.Category][]database.ServerTopic
}

func (m *TopicManager) Load() error {
	topics, err := database.LoadGlobalTopics()
	if err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())

	m.Topics = make(map[database.Category][]database.GlobalTopic)
	m.ServerTopics = make(map[int64]map[database.Category][]database.ServerTopic)

	for _, topic := range topics {
		if m.Topics[topic.Category] == nil {
			m.Topics[topic.Category] = make([]database.GlobalTopic, 0)
		}

		m.Topics[topic.Category] = append(m.Topics[topic.Category], topic)
	}

	return nil
}

func (m *TopicManager) populateGuildTopics(guild int64) error {
	if m.ServerTopics[guild] != nil {
		return nil
	}

	m.ServerTopics[guild] = make(map[database.Category][]database.ServerTopic)

	for _, category := range database.TopicCategories {
		m.ServerTopics[guild][category] = make([]database.ServerTopic, 0)
	}

	serverTopics, err := database.LoadServerTopics(guild)
	if err != nil {
		return err
	}

	for _, topic := range serverTopics {
		m.ServerTopics[guild][topic.Category] = append(m.ServerTopics[guild][topic.Category], topic)
	}

	return nil
}

func (m *TopicManager) GetRandomTopic(guild int64, category database.Category) (*Topic, error) {
	if m.ServerTopics[guild] == nil {
		m.populateGuildTopics(guild)
	}

	topicCount := len(m.Topics[category])
	guildTopicCount := len(m.ServerTopics[guild][category])
	total := topicCount + guildTopicCount

	if total == 0 {
		return &Topic{
			ID:    0,
			Topic: "No matches found :(",
		}, nil
	}

	val := rand.Intn(total)

	if val < topicCount {
		return &Topic{
			ID:    m.Topics[category][val].ID,
			Topic: m.Topics[category][val].Topic,
		}, nil
	}

	val -= topicCount
	return &Topic{
		ID:    m.ServerTopics[guild][category][val].ID,
		Topic: m.ServerTopics[guild][category][val].Topic,
	}, nil
}

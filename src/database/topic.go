package database

import (
	"encoding/json"
	"os"
	"github.com/rivo-gg/reviver-go/src/util"
	"github.com/sirupsen/logrus"
)

type Category string

const (
	CategoryTopic Category = "topic"
	CategoryFact  Category = "fact"
)

var TopicCategories = []Category{
	CategoryTopic,
	CategoryFact,
}

type GlobalTopic struct {
	ID       int64    `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	Category Category `gorm:"type:varchar(255);index"`
	Topic    string   `gorm:"type:text"`
}

type ServerTopic struct {
	ID       int64    `gorm:"primaryKey;autoIncrement;uniqueIndex"`
	Category Category `gorm:"type:varchar(255);index"`
	Topic    string   `gorm:"type:text"`
	GuildID  int64    `gorm:"index"`
	UserID   int64    `gorm:"index"`
}

func CreateGlobalTopic(category Category, data string) (*GlobalTopic, error) {
	topic := &GlobalTopic{
		Category: category,
		Topic:    data,
	}

	if err := DB.Create(topic).Error; err != nil {
		return nil, err
	}

	return topic, nil
}

func DeleteGlobalTopic(id int64) error {
	if err := DB.Where("id = ?", id).Delete(&GlobalTopic{}).Error; err != nil {
		return err
	}
	return nil
}

func LoadGlobalTopics() ([]GlobalTopic, error) {
	var topics []GlobalTopic
	if err := DB.Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func PopulateGlobalTopics() error {
	data, err := os.ReadFile("data/data.json")
	if err != nil {
		return err
	}

	var loadedTopics = make(map[string][]string)
	existingTopics := []string{}

	topics, err := LoadGlobalTopics()
	if err != nil {
		return err
	}

	for _, topic := range topics {
		existingTopics = append(existingTopics, topic.Topic)
	}

	if err := json.Unmarshal(data, &loadedTopics); err != nil {
		return err
	}

	for _, topic := range loadedTopics["topics"] {
		if util.Contains(existingTopics, topic) {
			continue
		}

		_, err := CreateGlobalTopic(CategoryTopic, topic)
		if err != nil {
			return err
		}

		logrus.Info("Added topic: ", topic)
	}

	for _, fact := range loadedTopics["facts"] {
		if util.Contains(existingTopics, fact) {
			continue
		}

		_, err := CreateGlobalTopic(CategoryFact, fact)
		if err != nil {
			return err
		}

		logrus.Info("Added fact: ", fact)
	}

	return nil
}

func CreateServerTopic(category Category, data string, guildID int64, userID int64) (*ServerTopic, error) {
	topic := &ServerTopic{
		Category: category,
		Topic:    data,
		GuildID:  guildID,
		UserID:   userID,
	}

	if err := DB.Create(topic).Error; err != nil {
		return nil, err
	}

	return topic, nil
}

func DeleteServerTopic(id int64) error {
	if err := DB.Where("id = ?", id).Delete(&ServerTopic{}).Error; err != nil {
		return err
	}
	return nil
}

func LoadServerTopics(guildID int64) ([]ServerTopic, error) {
	var topics []ServerTopic
	if err := DB.Where("guild_id = ?", guildID).Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

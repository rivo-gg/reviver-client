package database

import "strings"

type Guild struct {
	ID      int64  `gorm:"primaryKey;autoIncrement:false;uniqueIndex"`
	Name    string `gorm:"type:varchar(255)"`
	OwnerID int64  `gorm:"index"`
	Config  string `gorm:"type:text"`
	Banned  bool   `gorm:"index"`
	Flags   string `gorm:"type:text"`
}

func CreateGuild(id int64, name string, ownerID int64) (*Guild, error) {
	guild := &Guild{
		ID:      id,
		Name:    name,
		OwnerID: ownerID,
		Config:  "",
		Banned:  false,
	}

	if err := DB.Create(guild).Error; err != nil {
		return nil, err
	}

	return guild, nil
}

func GetGuild(id int64) (*Guild, error) {
	var guild Guild
	if err := DB.Where("id = ?", id).First(&guild).Error; err != nil {
		return nil, err
	}
	return &guild, nil
}

func UpdateGuild(id int64, name string, ownerID int64, config string, banned bool, flags []string) (*Guild, error) {
	guild, err := GetGuild(id)
	if err != nil {
		return nil, err
	}

	guild.Name = name
	guild.OwnerID = ownerID
	guild.Config = config
	guild.Banned = banned
	guild.Flags = strings.Join(flags, ";")
	if err := DB.Save(guild).Error; err != nil {
		return nil, err
	}

	return guild, nil
}

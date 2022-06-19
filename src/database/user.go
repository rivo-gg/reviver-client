package database

import "strings"

type User struct {
	ID     int64  `gorm:"primaryKey;autoIncrement:false;uniqueIndex"`
	Name   string `gorm:"type:varchar(255)"`
	Banned bool   `gorm:"index"`
	Flags  string `gorm:"type:text"`
}

func CreateUser(id int64, name string) (*User, error) {
	user := &User{
		ID:     id,
		Name:   name,
		Banned: false,
		Flags:  "",
	}

	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(id int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id int64, name string, banned bool, flags []string) (*User, error) {
	user, err := GetUser(id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Banned = banned
	user.Flags = strings.Join(flags, ";")
	if err := DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

package database

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() {
	godotenv.Load()

	dsn := os.Getenv("DB_DSN")

	logrus.Info("Connecting to database: ", dsn)

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")))
	if err != nil {
		panic(err)
	}

	logrus.Info("Connected to database. Automigrating...")

	db.AutoMigrate(
		&GlobalTopic{},
		&ServerTopic{},
		&Guild{},
		&User{},
	)

	logrus.Info("Automigration complete.")

	DB = db
}

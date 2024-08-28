package main

import (
	"os"
	"os/signal"

	"github.com/rivo-gg/reviver-go/src/database"
	"github.com/rivo-gg/reviver-go/src/discord"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	})

	database.GetDB()
}

func main() {
	logrus.Info("Starting Reviver...")

	database.PopulateGlobalTopics()

	err, s, errs := discord.StartDiscordClient()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to start Discord client.")
	}

	logrus.Info("Reviver started.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	for {
		select {
		case err := <-errs:
			logrus.WithError(err.Error).Error("Error in command: ", err.Command)
		case <-stop:
			logrus.Info("Stopping Reviver...")
			s.Shutdown()
			return
		}
	}
}

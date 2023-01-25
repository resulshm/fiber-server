package main

import (
	"fmt"
	"os"

	"github.com/ResulShamuhammedov/fiber-server/pkg/repository"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Couldn't load env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	for i := 0; i < 25; i++ {
		err = repository.AddVidesToDatabase(db)
		if err != nil {
			logrus.Fatalf("Failed to insert videos to database: %s", err.Error())
		}
		logrus.Info(fmt.Sprintf("added %d entries", 4000*(i+1)))
	}
}

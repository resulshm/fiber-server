package main

import (
	"os"

	"github.com/ResulShamuhammedov/fiber-server/pkg/handlers"
	"github.com/ResulShamuhammedov/fiber-server/pkg/repository"
	"github.com/ResulShamuhammedov/fiber-server/pkg/service"
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

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	app := handler.InitRoutes()

	logrus.Info("Listening on port :3000")

	err = app.Listen(":3000")
	if err != nil {
		logrus.Fatalf("Failed to listen: %s", err.Error())
	}
}

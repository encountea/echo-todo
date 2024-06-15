package main

import (
	"log"
	"os"

	"github.com/encountea/echo-todo/config"
	"github.com/encountea/echo-todo/internal/handler"
	"github.com/encountea/echo-todo/internal/repository"
	"github.com/encountea/echo-todo/internal/service"
	"github.com/encountea/echo-todo/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("Config error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
	})
	if err != nil {
		log.Fatalf("Failed to initialize config: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := new(server.Server)

	log.Printf("Server is running on port: %v", cfg.ServerPort)
	if err := app.Run(cfg.ServerPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error to connect to server: %s", err.Error())
	}
}

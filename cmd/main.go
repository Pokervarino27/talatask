package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pokervarino27/talatask/cmd/routes"
	"github.com/pokervarino27/talatask/internal/infraestructure/config"
	logger "github.com/pokervarino27/talatask/internal/infraestructure/logger"
)

func main() {

	logger.Init()
	app := fiber.New()

	appConfig := config.NewAppConfig()

	routes.SetupRoutes(app, appConfig.Handler)

	logger.Info("Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

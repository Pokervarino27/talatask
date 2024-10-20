package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pokervarino27/talatask/internal/config"
)

func main() {

	appConfig := config.NewAppConfig()

	app := fiber.New()

	appConfig.Handler.SetupRoutes(app)

	log.Println("Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

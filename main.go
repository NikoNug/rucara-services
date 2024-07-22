package main

import (
	"log"
	"rucara-services/database"
	"rucara-services/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}

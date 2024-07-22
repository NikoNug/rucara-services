package routes

import (
	"rucara-services/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/register", controllers.UserSignUp)
	authGroup.Post("/login", controllers.UserLogin)
}

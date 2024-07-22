package routes

import (
	"rucara-services/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCommentRoutes(app *fiber.App) {
	commentGroup := app.Group("/comments")
	commentGroup.Post("/", controllers.CreateComment)
}

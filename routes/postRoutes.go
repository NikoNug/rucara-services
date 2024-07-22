package routes

import (
	"rucara-services/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App) {
	postGroup := app.Group("/posts")
	postGroup.Post("/", controllers.CreatePost)
	postGroup.Get("/", controllers.GetPosts)
	postGroup.Delete("/:id", controllers.DeletePost)
}

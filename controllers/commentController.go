package controllers

import (
	"rucara-services/database"
	"rucara-services/models"

	"github.com/gofiber/fiber/v2"
)

// CreateComment handles the creation of a new comment
func CreateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	_, err := database.DB.Exec("INSERT INTO comments (post_id, content) VALUES (?, ?)", comment.PostID, comment.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create comment"})
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}

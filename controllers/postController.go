package controllers

import (
	"rucara-services/database"
	"rucara-services/models"

	"github.com/gofiber/fiber/v2"
)

// CreatePost handles the creation of a new post
func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	_, err := database.DB.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", post.Title, post.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}

// GetPosts retrieves all posts from the database
func GetPosts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to scan post"})
		}
		posts = append(posts, post)
	}

	return c.JSON(posts)
}

// DeletePost handles the deletion of a post by its ID
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

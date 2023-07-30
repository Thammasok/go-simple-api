package post

import (
	postController "go-simple/app/post/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	// Post group
	postGroup := app.Group("/posts")

	// List All Posts
	postGroup.Get("/", postController.Index)

	// Show Post
	postGroup.Get("/:id", postController.Show)

	// Create Post
	postGroup.Post("/", postController.Create)

	// Update Post
	postGroup.Post("/:id", postController.Update)

	// Delete Post
	postGroup.Delete("/:id", postController.Delete)
}

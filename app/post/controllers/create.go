package post

import (
	"go-simple/app/post/models"
	"go-simple/database"

	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	post := models.Post{}
	if err := ctx.BodyParser(&post); err != nil {
		return err
	}

	if len(post.Content) < 1 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Content is required",
		})
	}

	if err := database.Connection().Create(&post).Error; err != nil {
		return err
	}

	return ctx.JSON(post)
}

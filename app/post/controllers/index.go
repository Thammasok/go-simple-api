package post

import (
	"go-simple/app/post/models"
	"go-simple/database"

	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	posts := []models.Post{}
	database.Connection().Find(&posts)

	return ctx.JSON(posts)
}

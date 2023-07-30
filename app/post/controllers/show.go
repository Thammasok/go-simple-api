package post

import (
	"go-simple/app/post/models"
	"go-simple/database"

	"github.com/gofiber/fiber/v2"
)

func Show(ctx *fiber.Ctx) error {
	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}

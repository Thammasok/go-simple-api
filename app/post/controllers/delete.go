package post

import (
	"go-simple/app/post/models"
	"go-simple/database"

	"github.com/gofiber/fiber/v2"
)

func Delete(ctx *fiber.Ctx) error {
	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = database.Connection().Delete(&post).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}

package migrations

import (
	"go-simple/app/post/models"
	"go-simple/database"
)

func Migrate() {
	database.Connection().AutoMigrate(&models.Post{})
}

package hello

import (
	helloControllers "go-simple/app/hello/controllers"

	"github.com/gofiber/fiber/v2"
)

func HelloRoutes(app *fiber.App) {
	// Post group
	helloGroup := app.Group("/v1/hello")

	helloGroup.Get("/", helloControllers.Hello)
}

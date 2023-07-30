package hello

import (
	helloControllers "go-simple/app/hello/controllers"

	"github.com/gofiber/fiber/v2"
)

func HelloRoutes(app *fiber.App) {
	app.Get("/", helloControllers.Hello)
}

package server

import (
	"fmt"
	"log"

	"go-simple/middleware"

	routeHello "go-simple/app/hello"
	routePost "go-simple/app/post"

	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) {
	// Create Fiber App
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: ErrorHandler,
	})

	// Initialize Logger
	file := middleware.Logger(app)
	defer file.Close()

	// Middlewares
	app.Use(middleware.Example)

	// Mount routes
	routeHello.HelloRoutes(app)
	routePost.PostRoutes(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

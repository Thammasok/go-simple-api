package main

import (
	"os"
	"strconv"

	"go-simple/migrations"
	"go-simple/server"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if !fiber.IsChild() {
		// Application initialization
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	server.CreateServer(port)
}

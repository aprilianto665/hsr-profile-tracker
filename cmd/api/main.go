package main

import (
	"hsr-profile-tracker/internal/database"
	"hsr-profile-tracker/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectRedis()

	app := fiber.New()
	app.Use(cors.New())

	routes.ProfileRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

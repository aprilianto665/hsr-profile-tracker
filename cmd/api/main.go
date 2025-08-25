package main

import (
	"hsr-profile-tracker/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main () {
	app := fiber.New()

	routes.ProfileRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
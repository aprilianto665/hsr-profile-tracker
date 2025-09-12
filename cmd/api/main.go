package main

import (
	"hsr-profile-tracker/internal/configs"
	"hsr-profile-tracker/internal/database"
	"hsr-profile-tracker/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectRedis()

	app := fiber.New()
	app.Use(cors.New())

	chars, err := configs.LoadCharacterWeights("internal/configs/character_weights.json")
	if err != nil {
		log.Fatal("Failed to load character weights:", err)
	}

	stats, err := configs.LoadStatWeights("internal/configs/stat_weights.json")
	if err != nil {
		log.Fatal("Failed to load effective stats:", err)
	}

	configs.CharacterWeights = chars
	configs.StatWeights = stats

	routes.ProfileRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"
	}

	log.Fatal(app.Listen(port))
}

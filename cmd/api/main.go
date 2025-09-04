package main

import (
	"hsr-profile-tracker/internal/configs"
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

	chars, err := configs.LoadCharacterWeights("internal/configs/character_weights.json")
	if err != nil {
		log.Fatal("Failed to load character weights:", err)
	}

	stats, err := configs.LoadEffectiveStats("internal/configs/effective_stats.json")
	if err != nil {
		log.Fatal("Failed to load effective stats:", err)
	}

	configs.EffectiveStats = stats
	configs.CharacterWeights = chars

	routes.ProfileRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

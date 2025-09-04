package main

import (
	"hsr-profile-tracker/internal/configs"
	"hsr-profile-tracker/internal/database"
	"hsr-profile-tracker/internal/model"
	"hsr-profile-tracker/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectRedis()

	app := fiber.New()
	app.Use(cors.New())

	stats, err := configs.LoadWeights[*model.EffectiveStats]("internal/configs/effective_stats.json")
	if err != nil {
		log.Fatal("Failed to load effective stats:", err)
	}

	chars, _ := configs.LoadWeights[*model.CharacterWeights]("internal/configs/character_weights.json")
	if err != nil {
		log.Fatal("Failed to load character weights:", err)
	}

	configs.EffectiveStats = stats
	configs.CharacterWeights = chars

	routes.ProfileRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

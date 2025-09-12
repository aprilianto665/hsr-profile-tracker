package main

import (
	"hsr-profile-tracker/internal/configs"
	"hsr-profile-tracker/internal/database"
	"hsr-profile-tracker/internal/routes"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	rstorage "github.com/gofiber/storage/redis/v2"
)

func main() {
	database.ConnectRedis()

	app := fiber.New()
	app.Use(cors.New())

	redisConfig := database.LoadRedisConfig()

	store := rstorage.New(rstorage.Config{
		Host:     redisConfig.Host,
		Port:     redisConfig.Port,
		Password: redisConfig.Password,
		Database: redisConfig.DB,
	})

	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: time.Minute,
		Storage:    store,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":  "error",
				"message": "Rate limit exceeded. Please try again later.",
			})
		},
	}))

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

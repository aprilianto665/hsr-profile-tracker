package database

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func LoadRedisConfig() RedisConfig {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	portStr := os.Getenv("REDIS_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 6379
	}

	password := os.Getenv("REDIS_PASSWORD")

	dbStr := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(dbStr)
	if err != nil || db < 0 {
		db = 0
	}

	return RedisConfig{Host: host, Port: port, Password: password, DB: db}
}

var Ctx = context.Background()
var Rdb *redis.Client

func ConnectRedis() {
	config := LoadRedisConfig()

	addr := config.Host + ":" + strconv.Itoa(config.Port)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
}

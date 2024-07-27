package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	Port        int
	Database    struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  time.Duration
	}
}

func Load() (*Config, error) {

	var cfg Config

	cfg.Environment = getEnv("ENVIRONMENT", "development")

	if cfg.Environment == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file")
		}
	}

	cfg.Port = getEnvAsInt("SERVER_PORT", 8080)
	cfg.Database.Dsn = getEnv("DATABASE_URL", "")
	cfg.Database.MaxIdleConns = getEnvAsInt("DATABASE_MAX_IDLE_CONNS", 25)
	cfg.Database.MaxOpenConns = getEnvAsInt("DATABASE_MAX_OPEN_CONNS", 25)
	cfg.Database.MaxIdleTime = time.Duration(getEnvAsInt("DATABASE_MAX_IDLE_TIME", 5000))

	return &cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

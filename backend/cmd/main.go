package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/Xav147/flashcards_backend/internal/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Could not load .env file")
	}
	cfg := config{
		addr: envOrDefault("ADDR", ":8080"),
		db: dbConfig{
			dsn: envOrDefault("MONGODB_URI", "mongodb://localhost:27017"),
		},
	}

	db, err := mongodb.ConnectMongoDb(cfg.db.dsn)
	if err != nil {
		log.Printf("failed to connect to mongodb: %s", err)
		os.Exit(1)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := db.Disconnect(ctx); err != nil {
			log.Printf("failed to disconnect mongodb client: %s", err)
		}
	}()

	api := application{
		config: cfg,
		db:     db,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server has failed to start with error: %s", err)
		os.Exit(1)
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

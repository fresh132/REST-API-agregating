package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatalln("DATABASE_URL not found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalln("Error create pool db:", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalln("db not found:", err)
	}

	log.Println("Connected to db")

	return pool
}

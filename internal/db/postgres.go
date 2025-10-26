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
		log.Fatalln("DATABASE_URL не задан в .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalln("Ошибка при создании пула подключений:", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalln("Не удалось подключиться к базе данных:", err)
	}

	log.Println("Подключение к базе данных установлено")

	return pool
}

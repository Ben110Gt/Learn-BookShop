package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var PgxDB *pgxpool.Pool

func ConnectPGX() {
	// load .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Warning: no .env file found, using environment variables")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// Build connection string
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s&timezone=%s",
		user, password, host, port, dbname, sslmode, timezone,
	)

	// Create connection pool
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Test connection
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Cannot ping database: %v\n", err)
	}

	log.Println("Connected to PostgreSQL via pgx!")

	PgxDB = pool
}

// Get connection
func GetPGX() *pgxpool.Pool {
	return PgxDB
}

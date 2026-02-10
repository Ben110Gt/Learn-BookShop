package server

import (
	"book/shop/internal/configs"
	"book/shop/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Server() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file Server")
	}
	app := fiber.New()

	configs.ConnectDatabase()

	// --- Routes ---
	routes.AuthRoutes(app)
	routes.Use5Routes(app)
	routes.BookRoutes(app)

	port := os.Getenv("APP_PORT")

	app.Listen(":" + port)
}

package routes

import (
	"book/shop/internal/configs"
	"book/shop/internal/handlers"
	"book/shop/internal/repository"
	"book/shop/internal/service"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	configs.ConnectPGX()

	db := configs.GetPGX()

	userRepo := repository.NewUserRepository(db)    // PGX or GORM
	userService := service.NewUserService(userRepo) // Service Layer
	userHandler := handlers.NewAuthHandler(userService)

	// --- Public Routes ---
	api := app.Group("/api")
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)
}

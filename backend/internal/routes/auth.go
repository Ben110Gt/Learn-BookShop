package routes

import (
	"book/shop/internal/configs"
	"book/shop/internal/handlers"
	"book/shop/internal/repository"
	"book/shop/internal/service"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	configs.ConnectDatabase()

	db := configs.GetDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewAuthHandler(userService)

	// --- Public Routes ---
	api := app.Group("BookShop/auth")
	api.Post("/register", userHandler.Register) //✅
	api.Post("/login", userHandler.Login)       //✅
}

package routes

import (
	"book/shop/internal/configs"
	"book/shop/internal/handlers"
	"book/shop/internal/middlewares"
	"book/shop/internal/repository"
	"book/shop/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Use5Routes(app *fiber.App) {
	configs.ConnectPGX()
	db := configs.GetPGX()

	userRepo := repository.NewUserRepository(db)         // PGX or GORM
	userService := service.NewUserService(userRepo)      // Service Layer
	userHandler := handlers.NewUserHandler(userService)

	auth := app.Group("/BookShop")
	auth.Use(middlewares.JWTMiddleware())

	// User
	user := auth.Group("/user")
	user.Use(middlewares.RoleMiddleware("user"))
	user.Get("/profile", userHandler.GetProfile)
	// Admin
	admin := auth.Group("/admin")
	admin.Use(middlewares.RoleMiddleware("admin"))
	admin.Get("/users", userHandler.GetAllUsers)
	admin.Delete("/user/:id", userHandler.DeleteUser)
	admin.Put("/user/:id", userHandler.UpdateUser)

}

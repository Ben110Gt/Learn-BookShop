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
	configs.ConnectDatabase()
	db := configs.GetDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	auth := app.Group("/")
	auth.Use(middlewares.JWTMiddleware())

	// User
	user := auth.Group("/customer")
	user.Use(middlewares.RoleMiddleware("customer"))
	user.Get("/profile", userHandler.GetProfile)
	// Admin
	admin := auth.Group("/admin")
	admin.Use(middlewares.RoleMiddleware("admin"))
	admin.Get("/users", userHandler.GetAllUsers)           //✅
	admin.Delete("/user/:user_id", userHandler.DeleteUser) // ✅
	admin.Put("/user/:user_id", userHandler.UpdateUser)
	admin.Get("/user/:user_id", userHandler.GetUserByID) // ✅

}

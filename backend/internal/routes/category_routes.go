package routes

import (
	"book/shop/internal/configs"
	"book/shop/internal/handlers"
	"book/shop/internal/middlewares"
	"book/shop/internal/repository"
	"book/shop/internal/service"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()
	categoryrepo := repository.NewCategoryRepository(db)
	categoryservice := service.NewCategoryService(categoryrepo)
	categoryhandler := handlers.NewCategoryHandler(categoryservice)
	auth := app.Group("/")
	auth.Use(middlewares.JWTMiddleware())
	// Admin
	admin := auth.Group("/admin")
	admin.Use(middlewares.RoleMiddleware("admin"))
	admin.Post("/category", categoryhandler.CreateCategory)
	admin.Delete("/category/:category_id", categoryhandler.DeleteCategory)
	admin.Get("/categories", categoryhandler.GetAllCategories)

}

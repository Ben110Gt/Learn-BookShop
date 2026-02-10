package routes

import (
	"book/shop/internal/configs"
	"book/shop/internal/handlers"
	"book/shop/internal/middlewares"
	"book/shop/internal/repository"
	"book/shop/internal/service"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	configs.ConnectDatabase()
	db := configs.GetDB()

	bookrepo := repository.NewBookRepository(db)
	bookservice := service.NewBookService(bookrepo)
	bookhandler := handlers.NewBookHandler(bookservice)

	auth := app.Group("/")
	auth.Use(middlewares.JWTMiddleware())

	// Admin
	admin := auth.Group("/admin")
	admin.Use(middlewares.RoleMiddleware("admin"))
	admin.Post("/book", bookhandler.CreateBook)
	admin.Get("/book/:book_id", bookhandler.GetBookByID)
	admin.Get("/books", bookhandler.GetAllBooks)
	admin.Delete("/book/:book_id", bookhandler.DeleteBook)

}

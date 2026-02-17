package handlers

import (
	"book/shop/internal/domain/book"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	req := new(book.CreateBookRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	book, err := h.service.CreateBook(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(book)
}

// GetBookByID
func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("book_id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid book ID",
		})
	}
	book, err := h.service.GetBookByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(book)
}

// GetAllBooks
func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.service.GetAllBooks(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(books)
}

// DeleteBook
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("book_id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid book ID",
		})
	}
	err := h.service.DeleteBook(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}

// UpdateBook
// func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
// 	id := c.Params("book_id")
// 	if id == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "invalid book ID",
// 		})
// 	}
// 	req := new(book.Book)
// 	if err := c.BodyParser(req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "cannot parse JSON",
// 		})
// 	}
// 	req.ID = id
// 	updatedBook, err := h.service.UpdateBook(c.Context(), req)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	return c.Status(200).JSON(updatedBook)
// }

package handlers

import (
	"book/shop/internal/domain/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{service: service}

}

// GetProfile
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid or missing user_id in context",
		})
	}
	user, err := h.service.GetUserByID(c.Context(), userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(user)
}

// GetAllUsers
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUser(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(users)
}

// DeleteUser
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("user_id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}
	err := h.service.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

// UpdateUser
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}
	req := new(user.UpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	err = h.service.UpdateUser(c.Context(), strconv.Itoa(id), *req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

// GetUserByID
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("user_id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}
	user, err := h.service.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(user)
}

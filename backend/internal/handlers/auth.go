package handlers

import (
	"book/shop/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Auth user.Service
}

func NewAuthHandler(service user.Service) *AuthHandler {
	return &AuthHandler{Auth: service}
}

// Register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	req := new(user.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	u, err := h.Auth.Register(c.Context(), *req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	u.Password = ""

	return c.Status(201).JSON(u)

}

// Login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(user.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}
	res, err := h.Auth.Login(c.Context(), *req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	LoginResponse := &user.LoginResponse{
		Token:    res.Token,
		UserID:   res.UserID,
		Username: res.Username,
		Role:     res.Role,
	}
	return c.Status(200).JSON(LoginResponse)

}

// GetProfile

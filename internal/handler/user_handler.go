package handler

import (
	"app/internal/responses"
	"app/internal/services"
	t "app/internal/types"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		UserService: userService,
	}
}

func (uh *userHandler) MyInfo(c *fiber.Ctx) error {
	tokenPayload, ok := c.Locals("TokenPayload").(*t.TokenPayload)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	user, err := uh.UserService.GetUserProfile(c, tokenPayload.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get user")
	}
	if user == nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return responses.Ok(c, "User info retrieved successfully", user)
}

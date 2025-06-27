package utils

import (
	"app/internal/responses"
	"app/internal/validation"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if errorsMap := validation.CustomErrorMessages(err); len(errorsMap) > 0 {
		return responses.Error(c, fiber.StatusBadRequest, "Bad Request", errorsMap)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return responses.Error(c, fiberErr.Code, fiberErr.Message, nil)
	}

	return responses.Error(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
}

func NotFoundHandler(c *fiber.Ctx) error {
	return responses.Error(c, fiber.StatusNotFound, "Endpoint Not Found", nil)
}

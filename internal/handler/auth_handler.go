package handler

import (
	r "app/internal/responses"
	s "app/internal/services"
	valid "app/internal/validation"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	RefreshTokens(c *fiber.Ctx) error
}
type authHandler struct {
	AuthService  s.AuthService
	UserService  s.UserService
	TokenService s.TokenService
	// EmailService s.EmailService
}

func NewAuthHandler(authService s.AuthService, userService s.UserService, tokenService s.TokenService) AuthHandler {
	return &authHandler{
		AuthService:  authService,
		UserService:  userService,
		TokenService: tokenService,
		// EmailService: emailService,
	}
}

func (ah *authHandler) Register(c *fiber.Ctx) error {
	var req valid.Register
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(r.Error(c, fiber.StatusBadRequest, "Invalid request body", nil))
	}
	user, err := ah.UserService.CreateUser(c, &valid.CreateUser{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(r.Error(c, fiber.StatusInternalServerError, "Failed to create user", nil))
	}
	tokens, err := ah.TokenService.GenerateAuthTokens(c, user)
	if err != nil {
		return err
	}
	return r.Ok(c, "User registered successfully", tokens)
}

func (ah *authHandler) Login(c *fiber.Ctx) error {
	var req valid.Login
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := ah.AuthService.Login(c, &req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	tokens, err := ah.TokenService.GenerateAuthTokens(c, user)
	if err != nil {
		return err
	}
	return r.Ok(c, "Login successful", tokens)
}

func (ah *authHandler) RefreshTokens(c *fiber.Ctx) error {
	var req valid.RefreshToken

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	tokens, err := ah.AuthService.RefreshAuth(c, &req)
	if err != nil {
		return err
	}

	return r.Ok(c, "Tokens refreshed successfully", tokens)
}

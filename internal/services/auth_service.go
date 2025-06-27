package services

import (
	"app/internal/responses"
	"app/internal/validation"
	"app/pkg/database/ent"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthService interface {
	Register(c *fiber.Ctx, req *validation.Register) (*ent.User, error)
	Login(c *fiber.Ctx, req *validation.Login) (*ent.User, error)
	Logout(c *fiber.Ctx, req *validation.Logout) error
	RefreshAuth(c *fiber.Ctx, req *validation.RefreshToken) (*responses.Tokens, error)
	ResetPassword(c *fiber.Ctx, query *validation.Token, req *validation.UpdatePassOrVerify) error
	VerifyEmail(c *fiber.Ctx, query *validation.Token) error
}

// giống kiểu Dependency Injection
// AuthService struct implements the AuthService interface
type authService struct {
	Log      *logrus.Logger
	DB       *ent.Client
	Validate *validator.Validate
	// UserService  UserService
	// TokenService TokenService
}

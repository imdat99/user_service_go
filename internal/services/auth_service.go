package services

import (
	"app/internal/responses"
	t "app/internal/types"
	"app/internal/utils"
	"app/internal/validation"
	"app/pkg/database/ent"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthService interface {
	Register(c *fiber.Ctx, req *validation.Register) (*ent.User, error)
	Login(c *fiber.Ctx, req *validation.Login) (*ent.User, error)
	Logout(c *fiber.Ctx) error
	RefreshAuth(c *fiber.Ctx, req *validation.RefreshToken) (*responses.Tokens, error)
	// ResetPassword(c *fiber.Ctx, query *validation.Token, req *validation.UpdatePassOrVerify) error
	// VerifyEmail(c *fiber.Ctx) error
}

// giống kiểu Dependency Injection
// AuthService struct implements the AuthService interface
type authService struct {
	Log          *logrus.Logger
	DB           *ent.Client
	Validate     *validator.Validate
	UserService  UserService
	TokenService TokenService
}

func NewAuthService(db *ent.Client, va *validator.Validate, uS UserService, tS TokenService) AuthService {
	return &authService{
		Log:          utils.Log,
		DB:           db,
		Validate:     va,
		UserService:  uS,
		TokenService: tS,
	}
}

func (s *authService) Register(c *fiber.Ctx, req *validation.Register) (*ent.User, error) {
	if err := s.Validate.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.UserService.CreateUser(c, &validation.CreateUser{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	// tokens, err := s.TokenService.GenerateAuthTokens(c, user)
	// if err != nil {
	// 	return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to generate tokens")
	// }

	return user, nil
}

func (s *authService) Login(c *fiber.Ctx, req *validation.Login) (*ent.User, error) {
	if err := s.Validate.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.UserService.GetUserByEmail(c, req.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}
	user.Update().SetLastLoginAt(utils.GetCurrentTime()).Exec(c.Context())
	return user, nil
}
func (s *authService) Logout(c *fiber.Ctx) error {
	tokenPayload := c.Locals("TokenPayload").(*t.TokenPayload)
	// token, err := s.TokenService.GetTokenByUserID(c, req.RefreshToken)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusNotFound, "Token not found")
	// }
	if err := s.TokenService.DeleteToken(c, tokenPayload.Type, tokenPayload.UserID, tokenPayload.JTI); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to revoke refresh token")
	}

	return nil
}

func (s *authService) RefreshAuth(c *fiber.Ctx, req *validation.RefreshToken) (*responses.Tokens, error) {
	if err := s.Validate.Struct(req); err != nil {
		return nil, err
	}

	token, err := s.TokenService.GetTokenByUserID(c, req.RefreshToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
	}

	user, err := s.UserService.GetUserByID(c, token.UserID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
	}

	if err := s.TokenService.DeleteToken(c, token.TokenType, token.UserID, token.Jti); err != nil {
		return nil, fiber.ErrInternalServerError
	}
	newTokens, err := s.TokenService.GenerateAuthTokens(c, user)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return newTokens, err
}

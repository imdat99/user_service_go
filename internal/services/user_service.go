package services

import (
	"app/internal/utils"
	"app/internal/validation"
	"app/pkg/database/ent"
	userEntity "app/pkg/database/ent/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(c *fiber.Ctx, req *validation.CreateUser) (*ent.User, error)
	GetUserByID(c *fiber.Ctx, id string) (*ent.User, error)
	GetUserByEmail(c *fiber.Ctx, email string) (*ent.User, error)
	GetUserProfile(c *fiber.Ctx, id string) (*ent.User, error)
	// UpdateUser(id string, user *ent.User) error
	// DeleteUser(id string) error
}

type userService struct {
	Log      *logrus.Logger
	Ent      *ent.Client
	Validate *validator.Validate
}

func NewUserService(Ent *ent.Client, validate *validator.Validate) UserService {
	return &userService{
		Log:      logrus.New(),
		Ent:      Ent,
		Validate: validate,
	}
}
func (s *userService) CreateUser(c *fiber.Ctx, req *validation.CreateUser) (*ent.User, error) {
	if err := s.Validate.Struct(req); err != nil {
		return nil, err
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		s.Log.Errorf("Failed hash password: %+v", err)
		return nil, err
	}
	user, createErr := s.Ent.User.Create().
		SetEmail(req.Email).
		SetFirstName("").
		SetLastName("").
		SetPasswordHash(hashedPassword).Save(c.Context())

	if createErr != nil {
		s.Log.Errorf("Failed to create user: %+v", createErr)
		return nil, createErr
	}
	return user, nil
}
func (s *userService) GetUserByID(c *fiber.Ctx, id string) (*ent.User, error) {
	user, err := s.Ent.User.Get(c.Context(), id)
	if err != nil {
		s.Log.Errorf("Failed to get user by ID: %+v", err)
		return nil, err
	}
	user.PasswordHash = "" // Clear password hash before returning
	return user, nil
}

func (s *userService) GetUserByEmail(c *fiber.Ctx, email string) (*ent.User, error) {
	user, err := s.Ent.User.Query().Where(userEntity.EmailEQ(email)).Only(c.Context())
	if err != nil {
		s.Log.Errorf("Failed to get user by email: %+v", err)
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserProfile(c *fiber.Ctx, id string) (*ent.User, error) {
	user, err := s.Ent.User.Query().
		Where(userEntity.ID(id)).
		WithUserProfile().
		Only(c.Context())
	if err != nil {
		s.Log.Errorf("Failed to get user profile: %+v", err)
		return nil, err
	}
	return user, nil
}

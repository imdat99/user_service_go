package services

import (
	"app/internal/config"
	t "app/internal/types"
	"app/internal/utils"
	database "app/pkg/database"
	"app/pkg/database/ent"
	"fmt"
	"time"

	res "app/internal/responses"

	userTokenEntity "app/pkg/database/ent/usertoken"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type TokenService interface {
	// GenerateToken(payload t.TokenPayload) (string, error)
	saveToken(c *fiber.Ctx, token string, tokenPayload t.TokenPayload) error
	DeleteToken(c *fiber.Ctx, tokenType userTokenEntity.TokenType, userID string, jti string) error
	DeleteAllToken(c *fiber.Ctx, userID string) error
	GetTokenByUserID(c *fiber.Ctx, tokenStr string) (*ent.UserToken, error)
	GenerateAuthTokens(c *fiber.Ctx, user *ent.User) (*res.Tokens, error)
	// GenerateResetPasswordToken(c *fiber.Ctx, req *validation.ForgotPassword) (string, error)
	// GenerateVerifyEmailToken(c *fiber.Ctx, user *ent.User) (*string, error)
}

type tokenService struct {
	Log      *logrus.Logger
	DB       *database.DBClient
	Validate *validator.Validate
}

func NewTokenService(db *database.DBClient, validate *validator.Validate) TokenService {
	return &tokenService{
		Log:      utils.Log,
		DB:       db,
		Validate: validate,
	}
}

func generateToken(payload t.TokenPayload) (string, error) {
	claims := jwt.MapClaims{
		"sub":  payload.UserID,
		"iat":  payload.IAT,
		"exp":  payload.Exp,
		"type": payload.Type.String(),
		"jti":  payload.JTI,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWTSecret))
}

func (s *tokenService) saveToken(c *fiber.Ctx, token string, tokenPayload t.TokenPayload) error {
	if err := s.DeleteToken(c, tokenPayload.Type, tokenPayload.UserID, tokenPayload.JTI); err != nil {
		return err
	}
	if tokenPayload.Type == userTokenEntity.TokenTypeAccessToken {
		return s.DB.Redis.Set(c.Context(), fmt.Sprintf("%s:%s", tokenPayload.UserID, tokenPayload.JTI), token, time.Unix(tokenPayload.Exp, 0).Sub(time.Now().UTC())).Err()
	}
	tokenDoc := s.DB.Ent.UserToken.Create().
		SetUserID(tokenPayload.UserID).
		SetTokenType(tokenPayload.Type).
		SetTokenHash(utils.HashToken(token)).
		SetJti(tokenPayload.JTI).
		SetExpiresAt(time.Unix(tokenPayload.Exp, 0)).
		SetIPAddress(c.IP()).
		SetUserAgent(c.Get("User-Agent")).
		SetDeviceInfo(c.Get("Device-Info")).
		SetCreatedAt(time.Now().UTC())
	_, err := tokenDoc.Save(c.Context())
	if err != nil {
		s.Log.Errorf("Failed save token: %+v", err)
	}

	return err
}

func (s *tokenService) DeleteToken(c *fiber.Ctx, tokenType userTokenEntity.TokenType, userID string, jti string) error {
	if tokenType == userTokenEntity.TokenTypeAccessToken {
		err := s.DB.Redis.Del(c.Context(), fmt.Sprintf("%s:%s", userID, jti)).Err()
		if err != nil {
			s.Log.Errorf("Failed to delete access token from Redis: %+v", err)
		}
		return err
	}
	_, err := s.DB.Ent.UserToken.Delete().Where(userTokenEntity.TokenTypeEQ(tokenType)).
		Where(userTokenEntity.UserIDEQ(userID)).Where(userTokenEntity.JtiEQ(jti)).Exec(c.Context())
	if err != nil {
		s.Log.Errorf("Failed to delete token: %+v", err)
	}

	return err
}

func (s *tokenService) DeleteAllToken(c *fiber.Ctx, userID string) error {
	reErr := s.DB.Redis.Del(c.Context(), fmt.Sprintf("%s:*", userID)).Err()
	// result := s.DB.WithContext(c.Context()).Where("user_id = ?", userID).Delete(tokenDoc)
	_, err := s.DB.Ent.UserToken.Delete().Where(userTokenEntity.UserIDEQ(userID)).Exec(c.Context())

	if err != nil {
		s.Log.Errorf("Failed to delete all token: %+v", err)
	}
	if reErr != nil {
		s.Log.Errorf("Failed to delete all token from Redis: %+v", reErr)
		return reErr
	}

	return err
}

func (s *tokenService) GetTokenByUserID(c *fiber.Ctx, tokenStr string) (*ent.UserToken, error) {
	payload, err := utils.VerifyToken(tokenStr, config.JWTSecret, userTokenEntity.TokenTypeRefreshToken)
	if err != nil {
		return nil, err
	}
	result, err := s.DB.Ent.UserToken.Query().
		Where(userTokenEntity.UserIDEQ(payload.UserID)).
		Where(userTokenEntity.JtiEQ(payload.JTI)).First(c.Context())
	if err != nil {
		s.Log.Errorf("Failed get token by user id: %+v", err)
		return nil, err
	}

	return result, nil
}

func (s *tokenService) GenerateAuthTokens(c *fiber.Ctx, user *ent.User) (*res.Tokens, error) {
	jti := utils.GenerateJTI() // Generate a unique identifier for the token
	accessTokenExpires := time.Now().UTC().Add(time.Minute * time.Duration(config.JWTAccessExp)).Unix()
	tokenPayload := &t.TokenPayload{
		UserID: user.ID,
		IAT:    time.Now().Unix(),
		Exp:    accessTokenExpires,
		Type:   userTokenEntity.TokenTypeAccessToken,
		JTI:    jti,
	}
	accessToken, err := generateToken(*tokenPayload)
	if err != nil {
		s.Log.Errorf("Failed generate token: %+v", err)
		return nil, err
	}

	tokenPayload.Type = userTokenEntity.TokenTypeRefreshToken

	tokenPayload.Exp = time.Now().UTC().Add(time.Hour * 24 * time.Duration(config.JWTRefreshExp)).Unix()
	refreshToken, err := generateToken(*tokenPayload)
	if err != nil {
		s.Log.Errorf("Failed generate token: %+v", err)
		return nil, err
	}

	if err = s.saveToken(c, refreshToken, *tokenPayload); err != nil {
		return nil, err
	}
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: false,
		Path:     "/",
		Domain:   "",
		SameSite: fiber.CookieSameSiteLaxMode, // Thử dùng Lax thay vì None
	})
	return &res.Tokens{
		Access: res.TokenExpires{
			Token:   accessToken,
			Expires: accessTokenExpires,
		},
		Refresh: res.TokenExpires{
			Token:   refreshToken,
			Expires: tokenPayload.Exp,
		},
	}, nil
}

// func (s *tokenService) GenerateResetPasswordToken(c *fiber.Ctx, req *validation.ForgotPassword) (string, error) {
// 	if err := s.Validate.Struct(req); err != nil {
// 		return "", err
// 	}

// 	user, err := s.UserService.GetUserByEmail(c, req.Email)
// 	if err != nil {
// 		return "", err
// 	}

// 	expires := time.Now().UTC().Add(time.Minute * time.Duration(config.JWTResetPasswordExp))
// 	resetPasswordToken, err := s.GenerateToken(user.ID.String(), expires, config.TokenTypeResetPassword)
// 	if err != nil {
// 		s.Log.Errorf("Failed generate token: %+v", err)
// 		return "", err
// 	}

// 	if err = s.SaveToken(c, resetPasswordToken, user.ID.String(), config.TokenTypeResetPassword, expires); err != nil {
// 		return "", err
// 	}

// 	return resetPasswordToken, nil
// }

// func (s *tokenService) GenerateVerifyEmailToken(c *fiber.Ctx, user *model.User) (*string, error) {
// 	expires := time.Now().UTC().Add(time.Minute * time.Duration(config.JWTVerifyEmailExp))
// 	verifyEmailToken, err := s.GenerateToken(user.ID.String(), expires, config.TokenTypeVerifyEmail)
// 	if err != nil {
// 		s.Log.Errorf("Failed generate token: %+v", err)
// 		return nil, err
// 	}

// 	if err = s.SaveToken(c, verifyEmailToken, user.ID.String(), config.TokenTypeVerifyEmail, expires); err != nil {
// 		return nil, err
// 	}

// 	return &verifyEmailToken, nil
// }

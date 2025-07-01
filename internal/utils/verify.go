package utils

import (
	"app/internal/types"
	userTokenEntity "app/pkg/database/ent/usertoken"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenStr string, secret string, tokenType userTokenEntity.TokenType) (*types.TokenPayload, error) {
	token, err := jwt.Parse(tokenStr, func(_ *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	tokenPayload := &types.TokenPayload{}
	if err != nil || !token.Valid {
		return tokenPayload, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return tokenPayload, errors.New("invalid token claims")
	}

	jwtType, ok := claims["type"].(string)
	if !ok || jwtType != tokenType.String() {
		return tokenPayload, errors.New("invalid token type")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return tokenPayload, errors.New("invalid token sub")
	}
	jti, ok := claims["jti"].(string)
	if !ok {
		return tokenPayload, errors.New("invalid token jti")
	}
	return &types.TokenPayload{
		UserID: userID,
		JTI:    jti,
		Type:   tokenType,
		IAT:    int64(claims["iat"].(float64)),
		Exp:    int64(claims["exp"].(float64)),
	}, nil
}

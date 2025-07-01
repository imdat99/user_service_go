package types

import (
	userTokenEntity "app/pkg/database/ent/usertoken"
)

type TokenPayload struct {
	UserID string                    `json:"sub"`
	JTI    string                    `json:"jti"`
	Type   userTokenEntity.TokenType `json:"type"`
	IAT    int64                     `json:"iat"`
	Exp    int64                     `json:"exp"`
}

package dtos

import (
	"app/pkg/database/ent/user"
	"time"
)

type User struct {
	ID string `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"password_hash,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// DateOfBirth holds the value of the "date_of_birth" field.
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`
	// EmailVerified holds the value of the "email_verified" field.
	EmailVerified bool `json:"email_verified,omitempty"`
	// PhoneVerified holds the value of the "phone_verified" field.
	PhoneVerified bool `json:"phone_verified,omitempty"`
	// AccountStatus holds the value of the "account_status" field.
	AccountStatus user.AccountStatus `json:"account_status,omitempty"`
	// LastLoginAt holds the value of the "last_login_at" field.
	LastLoginAt time.Time `json:"last_login_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

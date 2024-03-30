package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"primary_key"`
	UserID       string    `json:"user_id" gorm:"unique"`
	FirstName    *string   `json:"first_name" valid:"required"`
	LastName     *string   `json:"last_name" valid:"required"`
	Password     *string   `json:"password" valid:"required"`
	Email        *string   `json:"email" valid:"required,email" gorm:"unique"`
	Phone        *string   `json:"phone" valid:"required" gorm:"unique"`
	AvatarUrl    *string   `json:"avatar_url"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// BeforeCreate hook to check validity of table fields
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

	if user.FirstName == nil {
		return errors.New("First name is required")
	}
	if user.LastName == nil {
		return errors.New("Last name is required")
	}
	if user.Email == nil {
		return errors.New("Email is required")
	}
	if user.Phone == nil {
		return errors.New("Phone is required")
	}
	if user.Password == nil {
		return errors.New("Password is required")
	}

	user.UserID = generateUniqueUserID()
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

// BeforeUpdate hook to automatically update Updated_at field
func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

// generateUniqueUserID generates a unique user ID
func generateUniqueUserID() string {
	// Implement a logic to generate a unique user ID (e.g., using UUID)
	return uuid.New().String()
}

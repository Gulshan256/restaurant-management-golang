package models

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        *string    `json:"name" validate:"required"`
	Description *string    `json:"description"`
	Price       *float64   `json:"price" validate:"required"`
	Food_image  *string    `json:"image"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Food_id     string    `json:"food_id"`
	Menu_id     *string    `json:"menu_id"`
}

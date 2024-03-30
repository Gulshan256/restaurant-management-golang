package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primary_key"`
	Menu_id     string    `json:"menu_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Category    string    `json:"category" validate:"required"`
	Description string    `json:"description"`
	Start_Date  *time.Time `json:"start_date"`
	End_Date    *time.Time `json:"end_date"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

package models

import (
	"gorm.io/gorm"

	"time"
)

// Note struct
type Note struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	NoteID    string    `json:"note_id" gorm:"unique"`
	NoteTitle *string   `json:"note_title"`
	NoteBody  *string   `json:"note_body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// table struct
type Table struct {
	gorm.Model
	Id             uint      `json:"id" gorm:"primary_key"`
	TableID        string    `json:"table_id" gorm:"unique"`
	NumberOfGuests *int      `json:"number_of_guests"`
	Table_number   *int64    `json:"table_number"`
	TableStatus    string    `json:"table_status"`
	TableCapacity  *int64    `json:"table_capacity"`
	TableType      string    `json:"table_type"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

// BeforeCreate hook to check validity of table fields
func (table *Table) BeforeCreate(tx *gorm.DB) (err error) {
	if table.Table_number == nil {
		return errors.New("Table number is required")
	}
	if table.TableCapacity == nil {
		return errors.New("Table capacity is required")
	}
	if table.TableType == "" {
		return errors.New("Table type is required")
	}
	if *table.NumberOfGuests <= 0 {
		return errors.New("Number of guests is required and must be greater than 0")
	}

	table.TableID = "TAB" + generateUniqueTableID()
	table.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

// BeforeUpdate hook to automatically update Updated_at field
func (table *Table) BeforeUpdate(tx *gorm.DB) (err error) {
	table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

// generateUniqueTableID generates a unique table ID
func generateUniqueTableID() string {
	// Implement a logic to generate a unique table ID (e.g., using UUID)
	return uuid.New().String()
}

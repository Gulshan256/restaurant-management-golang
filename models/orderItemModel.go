package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OrderItem struct
type OrderItem struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primary_key"`
	OrderItemID string    `json:"order_item_id" gorm:"unique"`
	Quantity    *int      `json:"quantity" gorm:"not null"`
	UnitPrice   *float64  `json:"unit_price" gorm:"not null"`
	FoodID      *string   `json:"food_id" gorm:"not null"`
	OrderID     string    `json:"order_id" gorm:"not null"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

// BeforeCreate hook to automatically generate OrderItemID
func (orderItem *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	orderItem.OrderItemID = "ORDITM-" + generateUniqueOrderItemID()
	orderItem.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

// BeforeUpdate hook to automatically update Updated_at field
func (orderItem *OrderItem) BeforeUpdate(tx *gorm.DB) (err error) {
	orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

// generateUniqueOrderItemID generates a unique order item ID
func generateUniqueOrderItemID() string {
	// Implement a logic to generate a unique order item ID (e.g., using UUID)
	return uuid.New().String()
}

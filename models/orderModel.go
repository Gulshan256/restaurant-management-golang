package models

import (
	"time"

	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primary_key"`
	OrderID     string    `json:"order_id" gorm:"unique"`
	TableID     *string   `json:"table_id" validate:"required"`
	OrderStatus string    `json:"order_status" gorm:"default:IN_PROGRESS" validate:"oneof=IN_PROGRESS COMPLETED CANCELLED"`
	OrderDate   time.Time `json:"order_date" validate:"required"`
}

// BeforeCreate hook to automatically generate OrderID
func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.OrderID = "ORD" + generateUniqueOrderItemID()
	order.OrderDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

// BeforeCreate hook to automatically OrderDate
func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	order.OrderDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return
}

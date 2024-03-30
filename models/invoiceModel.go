package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ID             uint      `json:"id" gorm:"primary_key"`
	InvoiceID      string    `json:"invoice_id"`
	OrderID        string    `json:"order_id"`
	PaymentMethod  *string    `json:"payment_method" validate:"oneof=CARD CASH UPI"`
	PaymentStatus  *string    `json:"payment_status" validate:"oneof=PAID UNPAID" gorm:"default:UNPAID"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	Tax            float64   `json:"tax"`
	Discount       float64   `json:"discount"`
	TotalPrice     float64   `json:"total_price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

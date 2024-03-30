package controllers

import (
	"net/http"
	"time"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type InvoiceViewFormat struct {
	InvoiceID      string      `json:"invoice_id"`
	OrderID        string      `json:"order_id"`
	PaymentMethod  string      `json:"payment_method"`
	PaymentStatus  *string     `json:"payment_status"`
	PaymentDue     interface{} `json:"payment_due"`
	PaymentDueDate time.Time   `json:"payment_due_date"`
	TableNumber    interface{} `json:"table_number"`
	Orderdetails   interface{} `json:"order_details"`
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {

		var invoice models.Invoice
		initializers.DB.Find(&invoice)
		c.JSON(200, gin.H{
			"message": "Get Orders",
			"status":  "success",
			"data":    invoice,
		})

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceID := c.Param("invoice_id") // Use "invoice_id" as parameter name

		var invoice models.Invoice

		if err := initializers.DB.Where("invoice_id = ?", invoiceID).First(&invoice).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Invoice not found",
					"status":  "failed",
					"error":   err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occurred while retrieving invoice",
			})
			return
		}

		var invoiceView InvoiceViewFormat
		allOrderitems, err := ItemsByOrder(invoiceID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error occurred while retrieving order items",
			})

			invoiceView.PaymentDueDate = invoice.PaymentDueDate

			invoiceView.PaymentMethod = "null"
			if invoice.PaymentMethod != nil {
				invoiceView.PaymentMethod = *invoice.PaymentMethod
			}

			invoiceView.InvoiceID = invoice.InvoiceID
			// invoiceView.PaymentStatus = *&invoice.PaymentStatus
			invoiceView.PaymentStatus = invoice.PaymentStatus
			invoiceView.PaymentDue = allOrderitems[0]["paymentDue"]
			invoiceView.TableNumber = allOrderitems[0]["tableNumber"]
			// invoiceView.OrderID = allOrderitems[0]["orderID"]
			invoiceView.Orderdetails = allOrderitems[0]["orderDetails"]

			c.JSON(http.StatusOK, gin.H{
				"message": "Invoice is available",
				"status":  "success",
				"data":    invoiceView,
			})

			c.JSON(http.StatusOK, gin.H{
				"message": "Invoice is available",
				"status":  "success",
				"data":    invoice,
			})
		}
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice models.Invoice
		var order models.Order

		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
				"status":  "failed",
				"error":   err.Error(),
			})
			return
		}

		validationError := validator.New().Struct(invoice)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"validationError": validationError.Error(),
			})
			return
		}

		if err := initializers.DB.Where("order_id = ?", invoice.OrderID).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Order not found",
				"status":  "failed",
				"error":   "Order not found",
			})
			return
		}

		invoice.PaymentDueDate = time.Now().AddDate(0, 0, 7)
		result := initializers.DB.Create(&invoice)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create invoice",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Invoice created successfully",
			"status":  "success",
			"data":    result,
		})
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice models.Invoice
		invoiceID := c.Param("invoice_id")

		if err := initializers.DB.Where("invoice_id = ?", invoiceID).First(&invoice).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Invoice not found",
				"status":  "failed",
				"error":   "Invoice not found",
			})
			return
		}

		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
				"status":  "failed",
				"error":   err.Error(),
			})
			return
		}

		staus := "UNPAID"
		if invoice.PaymentStatus == nil {
			invoice.PaymentStatus = &staus
		}

		result := initializers.DB.Save(&invoice)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update invoice",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Invoice updated successfully",
			"status":  "success",
			"data":    result,
		})

	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var invoice models.Invoice
		invoiceID := c.Param("invoice_id")

		if err := initializers.DB.Where("invoice_id = ?", invoiceID).First(&invoice).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Invoice not found",
				"status":  "failed",
				"error":   "Invoice not found",
			})
			return
		}

		result := initializers.DB.Delete(&invoice)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete invoice",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Invoice deleted successfully",
			"status":  "success",
			"data":    result,
		})
	}
}

func ItemsByOrder(invoiceID string) ([]map[string]interface{}, error) {
	var orderItems []models.OrderItem
	var orderItemsMap []map[string]interface{}

	if err := initializers.DB.Where("invoice_id = ?", invoiceID).Find(&orderItems).Error; err != nil {
		return nil, err
	}

	for _, orderItem := range orderItems {
		orderDetails := make(map[string]interface{})
		orderDetails["orderID"] = orderItem.OrderID
		orderItemsMap = append(orderItemsMap, orderDetails)
	}

	return orderItemsMap, nil
}

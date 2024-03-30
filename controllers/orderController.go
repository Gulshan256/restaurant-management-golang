package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/models"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {

		var order []models.Order
		initializers.DB.Find(&order)
		c.JSON(200, gin.H{
			"message": "Get Orders",
			"status":  "success",
			"data":    order,
		})

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id") // Use "order_id" as parameter name

		var order models.Order

		if err := initializers.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Order not found",
					"status":  "failed",
					"error":   "Order not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve order",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Order is available",
			"status":  "success",
			"data":    order,
		})
	}
}

func CreateOrder() gin.HandlerFunc {

	return func(c *gin.Context) {

		var order models.Order
		var table models.Table

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
				"status":  "failed",
				"error":   err.Error(),
			})

		}

		validationError := validator.New().Struct(order)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"validationError": validationError.Error(),
			})
			return
		}

		// order table is not nil
		if order.TableID != nil {
			if err := initializers.DB.Where("table_id = ?", *order.TableID).First(&table).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					c.JSON(http.StatusNotFound, gin.H{
						"message": "Table not found",
						"status":  "failed",
						"error":   "Table not found",
					})
					return
				}
			}
		}

		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := initializers.DB.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create order",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Order created",
			"status":  "success",
			"data":    order,
		})

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		var table models.Table
		var order models.Order

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if order.TableID != nil {
			if err := initializers.DB.First(&table, *order.TableID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					c.JSON(http.StatusNotFound, gin.H{
						"Message": "Table not found",
					})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to retrieve table",
				})
				return
			}
		}

		order.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err := initializers.DB.Model(&order).Where("order_id = ?", orderId).Updates(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update order",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Order updated",
			"status":  "success",
			"data":    order,
		})

	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")

		var order models.Order
		if err := initializers.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Order not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve order",
			})
			return
		}

		result := initializers.DB.Delete(&order)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to delete order",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Order deleted successfully",
			"status":  "success",
			"result":  result,
		})
	}
}

func OrderItemOrderCreator(order models.Order) string {

	order.OrderID = "ORD" + time.Now().Format("20060102150405")

	initializers.DB.Create(&order)
	return order.OrderID
}

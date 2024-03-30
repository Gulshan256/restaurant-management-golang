package controllers

import (
	"net/http"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderItempack struct {
	TableID    *string
	OrderItems []models.OrderItem
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

		var orderItem []models.OrderItem

		result := initializers.DB.Find(&orderItem)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get OrderItems",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Get OrderItems",
			"status":  "success",
			"data":    orderItem,
		})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemID := c.Param("order_item_id")

		var orderItem models.OrderItem

		result := initializers.DB.Where("order_item_id = ?", orderItemID).First(&orderItem)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get OrderItem",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Get OrderItem",
			"status":  "success",
			"data":    orderItem,
		})
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		var orderItempack OrderItempack
		var order models.Order

		if err := c.ShouldBindJSON(&orderItempack); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		orderItemsToBeInserted := make([]models.OrderItem, 0)

		order_id := OrderItemOrderCreator(order)
		for _, orderItem := range orderItempack.OrderItems {
			orderItem.OrderID = order_id

			validationsErr := validator.New().Struct(orderItem)
			if validationsErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": validationsErr.Error(),
				})
				return
			}

			var num = toFixed(*orderItem.UnitPrice, 2)
			orderItem.UnitPrice = &num

			orderItemsToBeInserted = append(orderItemsToBeInserted, orderItem)
		}

		result := initializers.DB.Create(orderItemsToBeInserted)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create OrderItem",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "OrderItem created",
			"status":  "success",
			"data":    orderItemsToBeInserted,
		})

	}
}

func GetOrderItemByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		var orderItem []models.OrderItem
		orderID := c.Param("order_id")

		result := initializers.DB.Where("order_id = ?", orderID).Find(&orderItem)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get OrderItems",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Get OrderItems",
			"status":  "success",
			"data":    orderItem,
		})

	}
}

// func ItemsByOrder(orderID string) ([]models.OrderItem, error) {
// 	var orderItems []models.OrderItem

// 	// Join OrderItem, Food, Order, and Table models
// 	err := initializers.DB.
// 		Joins("JOIN foods ON order_items.food_id = foods.id").
// 		Joins("JOIN orders ON order_items.order_id = orders.id").
// 		Joins("JOIN tables ON orders.table_id = tables.id").
// 		Where("orders.order_id = ?", orderID).
// 		Find(&orderItems).Error

// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to retrieve order items by order ID")
// 	}

// 	return orderItems, nil
// }

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItem models.OrderItem
		orderItemID := c.Param("order_item_id")

		if err := c.BindJSON(&orderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate required fields
		if orderItem.UnitPrice == nil || *orderItem.UnitPrice <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "UnitPrice is required and must be greater than zero"})
			return
		}

		if orderItem.Quantity == nil || *orderItem.Quantity <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity is required and must be greater than zero"})
			return
		}

		if orderItem.FoodID == nil || *orderItem.FoodID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FoodID is required"})
			return
		}

		// Fetch existing order item from the database
		existingOrderItem := models.OrderItem{}
		if err := initializers.DB.Where("order_item_id = ?", orderItemID).First(&existingOrderItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch existing order item"})
			return
		}

		// Update only specified fields
		if err := initializers.DB.Model(&existingOrderItem).Updates(map[string]interface{}{
			"UnitPrice": orderItem.UnitPrice,
			"Quantity":  orderItem.Quantity,
			"FoodID":    orderItem.FoodID,
		}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OrderItem updated",
			"status":  "success",
			"data":    existingOrderItem,
		})
	}
}

func DeleteOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Delete OrderItem",
		})
	}
}

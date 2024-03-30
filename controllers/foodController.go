package controllers

import (
	"math"
	"time"

	"net/http"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		var food []models.Food
		initializers.DB.Find(&food)
		c.JSON(200, gin.H{
			"data": food,
		})
	}
}

// GetFood retrieves food details by ID.
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodID := c.Param("food_id")

		var food models.Food
		if err := initializers.DB.Where("food_id = ?", foodID).First(&food).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Food not found",
					"status":  "failed",
					"error":   "Food not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve food",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Food is available",
			"status":  "success",
			"data":    food,
		})
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})

		}

		validationError := validator.New().Struct(food)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError,
			})
			return
		}

		// find menu id
		if err := initializers.DB.First(&menu, food.Menu_id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Menu not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve menu",
			})
			return
		}

		// create food
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Food_id = "food_" + time.Now().Format("20060102150405")
		var num = toFixed(*food.Price, 2)
		food.Price = &num
		result := initializers.DB.Create(&food)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create food",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Food created successfully",
			"status":  "success",
			"result":  result,
		})

	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output

}

func UpdateFood() gin.HandlerFunc {

	return func(c *gin.Context) {

		foodID := c.Param("food_id")
		var food models.Food
		if err := initializers.DB.Where("food_id = ?", foodID).First(&food).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Food not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve food",
			})
			return
		}

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		validationError := validator.New().Struct(food)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError,
			})
			return
		}

		// update food
		food.Updated_at = time.Now()
		result := initializers.DB.Save(&food)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to update food",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Food updated successfully",
			"status":  "success",
			"result":  result,
		})

	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodID := c.Param("food_id")

		var food models.Food
		if err := initializers.DB.Where("food_id = ?", foodID).First(&food).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Food not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve food",
			})
			return
		}

		result := initializers.DB.Delete(&food)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to delete food",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Food deleted successfully",
			"status":  "success",
			"result":  result,
		})
	}
}

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

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu []models.Menu
		initializers.DB.Find(&menu)

		if len(menu) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No menu found",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Get Menu",
			"status":  "success",
			"data":    menu,
		})
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuID := c.Param("menu_id")

		var menu models.Menu
		if err := initializers.DB.Where("menu_id = ?", menuID).First(&menu).Error; err != nil {
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

		c.JSON(http.StatusOK, gin.H{
			"message": "Menu is available",
			"status":  "success",
			"data":    menu,
		})
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Menu_id = "Menu_" + time.Now().Format("20060102150405")

		validationError := validator.New().Struct(menu)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError,
			})
			return
		}

		result := initializers.DB.Create(&menu)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create menu",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Menu created successfully",
			"status":  "success",
			"data":    menu,
		})

	}
}

func UpdateMenu() gin.HandlerFunc {

	return func(c *gin.Context) {

		menuID := c.Param("menu_id")
		var menu models.Menu
		if err := initializers.DB.Where("menu_id = ?", menuID).First(&menu).Error; err != nil {
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

		var updatedMenu models.Menu
		if err := c.BindJSON(&updatedMenu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
		}

		validationError := validator.New().Struct(updatedMenu)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationError,
			})
			return
		}

		updatedMenu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		result := initializers.DB.Model(&menu).Updates(updatedMenu)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to update menu",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Menu updated successfully",
			"status":  "success",
			"data":    updatedMenu,
		})

	}
}

func DeleteMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuID := c.Param("menu_id")

		var menu models.Menu
		if err := initializers.DB.Where("menu_id = ?", menuID).First(&menu).Error; err != nil {
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

		result := initializers.DB.Delete(&menu)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to delete menu",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Menu deleted successfully",
			"status":  "success",
		})
	}
}

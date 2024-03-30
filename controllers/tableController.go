package controllers

import (
	"net/http"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/models"
	"github.com/gin-gonic/gin"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {

		var tables []models.Table

		result := initializers.DB.Find(&tables)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to get Tables",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Get Tables",
			"status":  "success",
			"data":    tables,
		})

	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {

		tableID := c.Param("table_id")

		var table models.Table

		result := initializers.DB.First(&table, "table_id = ?", tableID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to get Table",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Get Table",
			"status":  "success",
			"data":    table,
		})

	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table

		if err := c.ShouldBindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request payload",
				"status":  "failed",
				"error":   err.Error(),
			})
			return
		}

		result := initializers.DB.Create(&table)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to create Table",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Table created",
			"status":  "success",
			"data":    table,
		})
	}
}

func UpdateTable() gin.HandlerFunc {

	return func(c *gin.Context) {
		var table models.Table
		tableID := c.Param("table_id")

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := initializers.DB.Model(&table).Where("table_id = ?", tableID).Updates(&table)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to update Table",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Table updated",
			"status":  "success",
			"data":    table,
		})
	}
}

func DeleteTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("table_id")

		var table models.Table

		result := initializers.DB.Where("table_id = ?", tableID).Delete(&table)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to delete Table",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Table deleted",
			"status":  "success",
		})
	}
}

package routes

import (
	"github.com/devGulshan/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(c *gin.Engine) {

	c.GET("/foods", controllers.GetFoods())
	c.GET("/foods/:id", controllers.GetFood())
	c.POST("/foods", controllers.CreateFood())
	c.PATCH("/foods/:id", controllers.UpdateFood())
	c.DELETE("/foods/:id", controllers.DeleteFood())
}

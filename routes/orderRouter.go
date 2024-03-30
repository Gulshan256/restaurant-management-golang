package routes

import (
	"github.com/devGulshan/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(c *gin.Engine) {

	c.GET("/orders", controllers.GetOrders())
	c.GET("/orders/:id", controllers.GetOrder())
	c.POST("/orders", controllers.CreateOrder())
	c.PATCH("/orders/:id", controllers.UpdateOrder())
	c.DELETE("/orders/:id", controllers.DeleteOrder())

}

package routes

import (
	"github.com/devGulshan/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(c *gin.Engine) {

	c.GET("/tables", controllers.GetTables())
	c.GET("/tables/:id", controllers.GetTable())
	c.POST("/tables", controllers.CreateTable())
	c.PATCH("/tables/:id", controllers.UpdateTable())
	c.DELETE("/tables/:id", controllers.DeleteTable())

}

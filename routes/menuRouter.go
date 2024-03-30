package routes

import (
	"github.com/devGulshan/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(c *gin.Engine) {

	c.GET("/menus", controllers.GetMenus())
	c.GET("/menu/:id", controllers.GetMenu())
	c.POST("/menu", controllers.CreateMenu())
	c.PATCH("/menu/:id", controllers.UpdateMenu())
	c.DELETE("/menu/:id", controllers.DeleteMenu())

}

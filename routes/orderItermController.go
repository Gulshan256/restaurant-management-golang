package routes



import (
	"github.com/devGulshan/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)



func OrderItemsRoutes(c *gin.Engine) {
	
	c.GET("/orderItems",controllers.GetOrderItems())
	c.GET("/orderItems/:id",controllers.GetOrderItem())
	c.GET("/orderItems-order/:id",controllers.GetOrderItemByOrder())
	c.POST("/orderItems",controllers.CreateOrderItem())
	c.PATCH("/orderItems/:id",controllers.UpdateOrderItem())
	c.DELETE("/orderItems/:id",controllers.DeleteOrderItem())

}
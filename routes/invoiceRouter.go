package routes

import (
	"github.com/devGulshan/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)




func InvoiceRoutes(c *gin.Engine) {

	c.GET("/invoices",controllers.GetInvoices())
	c.GET("/invoices/:id",controllers.GetInvoice())
	c.POST("/invoices",controllers.CreateInvoice())
	c.PATCH("/invoices/:id",controllers.UpdateInvoice())
	c.DELETE("/invoices/:id",controllers.DeleteInvoice())

}
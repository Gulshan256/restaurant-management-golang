package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/middleware"
	"github.com/devGulshan/restaurant-management/routes"
)

func init() {
	initializers.LoadEnvVariables() // Load the environment variables

	initializers.ConnectToDb() // Initialize the database connection
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	// router := gin.New()
	router.Use(gin.Logger())
	routes.UserRouter(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}

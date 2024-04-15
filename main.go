package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/middleware"
	"github.com/devGulshan/restaurant-management/routes"

	"github.com/devGulshan/restaurant-management/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func init() {
	initializers.LoadEnvVariables() // Load the environment variables

	initializers.ConnectToDb() // Initialize the database connection
}

// @title Restaurant Management API
// @version 1
// @description This is a simple restaurant management API
// @contact.name Gulshan Kumar
// @contact.url http://www.gulshankumar.com
// @contact.email https://www.gulshankumar.com
// @BasePath /api/v1
// @schemes http
// @host localhost:8080

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	// router := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // http://localhost:8080/swagger/index.html
	router.Use(gin.Logger())

	routes.UserRouter(router)

	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}

}

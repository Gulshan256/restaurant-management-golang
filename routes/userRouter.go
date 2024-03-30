package routes

import (
	"github.com/devGulshan/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)



func UserRouter(c *gin.Engine) {

	c.GET("/users",controllers.GetUsers())
	c.GET("/users/:id",controllers.GetUser())
	c.POST("/users/signup",controllers.SignUp())
	c.POST("/users/login",controllers.Login())


}
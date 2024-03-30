package middleware

import (
	"net/http"

	"github.com/devGulshan/restaurant-management/halpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Authorization token is required",
			})
			c.Abort()
			return
		}

		claims, _, err := halpers.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("phone", claims.Phone)
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}

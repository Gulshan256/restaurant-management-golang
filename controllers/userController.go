package controllers

import (
	"net/http"

	"github.com/devGulshan/restaurant-management/halpers"
	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user []models.User

		initializers.DB.Find(&user)

		c.JSON(200, gin.H{
			"message": "Get Users",
			"data":    user,
		})

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.Param("id")

		var user models.User
		result := initializers.DB.First(&user, "user_id = ?", userID)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to get User",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Get User",
			"data":    user,
		})

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		// convert the JSOn data comming from the request
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON",
				"status":  "failed",
				"error":   err,
			})
			return
		}

		// validate the data based on the user struct
		if validationError := validator.New().Struct(user); validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON",
				"status":  "failed",
				"error":   validationError,
			})
			return
		}

		// check if email is valid  and not already taken
		var existingUser models.User
		if err := initializers.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Email already taken",
				"status":  "failed",
				"error":   "Email already taken",
			})
			return
		}
		//check if phone is valid and not already taken by other user
		if err := initializers.DB.Where("phone = ?", user.Phone).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Phone already taken",
				"status":  "failed",
				"error":   "Phone already taken",
			})
			return
		}

		// hash password
		hashedPassword := HashPassword(*user.Password)

		user.Password = &hashedPassword

		// generate token and refresh token from helper function
		token, refreshToken, _ := halpers.GenrateAlltokens(*user.Email, *user.FirstName, *user.LastName, *user.Phone, user.UserID)
		user.Token = &token
		user.RefreshToken = &refreshToken

		// create user
		result := initializers.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to create User",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"data":    user,
		})

	}
}

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {

		var userLogin models.User

		// convert the JSOn data comming from the request
		if err := c.BindJSON(&userLogin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON",
				"status":  "failed",
				"error":   err,
			})
			return
		}

		// validate the data based on the user struct
		if validationError := validator.New().Struct(userLogin); validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON",
				"status":  "failed",
				"error":   validationError,
			})
			return
		}

		// check if email is valid  and not already taken
		var existingUser models.User
		if err := initializers.DB.Where("email = ?", userLogin.Email).First(&existingUser).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Email not found",
				"status":  "failed",
				"error":   "Email not found",
			})
			return
		}

		// verify password
		isValid, _ := VerifyPassword(*existingUser.Password, *userLogin.Password)
		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid password",
				"status":  "failed",
				"error":   "Invalid password ",
			})
			return
		}

		// generate token and refresh token from helper function
		token, refreshToken, _ := halpers.GenrateAlltokens(*existingUser.Email, *existingUser.FirstName, *existingUser.LastName, *existingUser.Phone, existingUser.UserID)

		// update token and refresh token
		existingUser.Token = &token
		existingUser.RefreshToken = &refreshToken

		// update user
		result := initializers.DB.Save(&existingUser)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error. Failed to update User",
				"status":  "failed",
				"error":   result.Error,
			})
			return
		}

		// return user
		c.JSON(http.StatusOK, gin.H{
			"message": "User logged in successfully",
			"data":    existingUser,
		})

	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update User",
		})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete User",
		})
	}
}

func HashPassword(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providePassword))
	if err != nil {
		return false, "Password does not match"
	}
	return true, "Password match"
}

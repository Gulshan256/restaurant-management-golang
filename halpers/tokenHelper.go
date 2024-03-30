package halpers

import (
	"os"
	"time"

	"github.com/devGulshan/restaurant-management/initializers"
	"github.com/devGulshan/restaurant-management/models"
	"github.com/golang-jwt/jwt"
)

type Signeddetails struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	UserID    string `json:"user_id"`
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

// halpers.GenrateAlltokens(*existingUser.Email, *existingUser.FirstName, *existingUser.LastName, *existingUser.Phone, user.UserID)
func GenrateAlltokens(email string, firstName string, lastName string, phone string, userID string) (string, string, error) {

	claims := &Signeddetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &Signeddetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "",
			"",
			err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "",
			"",
			err
	}

	return token, refreshToken, nil

}

func UpdateAllTokens(token string, refreshToken string, userId string) {

	// get the user with the given user id
	var existingUser models.User
	initializers.DB.Where("user_id = ?", userId).First(&existingUser)

	// update the token and refresh token
	initializers.DB.Model(&existingUser).Update("token", token)
	initializers.DB.Model(&existingUser).Update("refresh_token", refreshToken)

}

func ValidateToken(token string) (claims *Signeddetails, msg string, err error) {

	// parse the token
	tkn, err := jwt.ParseWithClaims(token, &Signeddetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	// check token is invalid
	if err != nil {
		return nil, "Invalid token", err
	}

	claims, ok := tkn.Claims.(*Signeddetails)
	if !ok {
		return nil, "Invalid token", err
	}

	// check token is expired
	if time.Now().Local().Unix() > claims.ExpiresAt {
		return nil, "Token expired", err
	}

	return claims, "Token is valid", nil

}

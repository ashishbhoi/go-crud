package middlewares

import (
	"errors"
	"github.com/ashishbhoi/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ValidateToken(tokenString string) (models.PublicUser, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	var user models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		err = nil
		// Check user in database
		models.DB.Where("id = ?", claims["id"]).First(&user)
		if user.ID == 0 {
			err = errors.New("invalid login, Please login again")
		}
	} else {
		err = errors.New("invalid login, Please login again")
	}
	return models.PublicUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, err
}

func AuthFilter(context *gin.Context) {
	// Get the token from the header
	cookie, err := context.Cookie("Authorization")
	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"error": "Please login"})
		return
	}
	// Validate the token
	claims, err := ValidateToken(cookie)

	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}
	context.Set("user", claims)

	context.Next()
}

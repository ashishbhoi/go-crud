package middlewares

import (
	"errors"
	"github.com/ashishbhoi/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
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

		// Check if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			err = errors.New("token expired")
		}

		// Check user in database

		models.DB.Where("id = ?", claims["id"]).First(&user)
		if user.ID == 0 {
			err = errors.New("invalid token")
		}

	} else {
		err = errors.New("invalid token")
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
	header := context.Request.Header.Get("Authorization")
	if header == "" {
		context.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing"})
		return
	}
	authHeader := strings.Split(context.Request.Header.Get("Authorization"), " ")
	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		context.AbortWithStatusJSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
		return
	}

	// Validate the token
	token := authHeader[1]
	claims, err := ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}
	context.Set("user", claims)

	context.Next()
}

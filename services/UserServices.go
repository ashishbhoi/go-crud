package services

import (
	"github.com/ashishbhoi/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

type userModel struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func jwtToken(user userModel) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"email":     user.Email,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckUser(email string) bool {
	var user models.User
	models.DB.Where("email = ?", email).First(&user)
	if user.Email == email {
		return true
	}
	return false
}

func CreateUser(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if CheckUser(user.Email) {
		context.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "User Already Exists"})
		return
	}
	models.DB.Create(&user)
	displayUser := userModel{user.ID, user.FirstName, user.LastName, user.Email}
	token, err := jwtToken(displayUser)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.SetSameSite(http.SameSiteNoneMode)
	context.SetCookie("Authorization", token, 3600*24*7, "", "", false, true)
	context.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully and Logged In"})
}

func VerifyUser(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var user models.User
	var verifyUser models.User
	err := context.BindJSON(&verifyUser)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Where("email = ?", verifyUser.Email).First(&user)
	if user.Email == verifyUser.Email && user.Password == verifyUser.Password {
		displayUser := userModel{user.ID, user.FirstName, user.LastName, user.Email}
		token, err := jwtToken(displayUser)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.SetSameSite(http.SameSiteNoneMode)
		context.SetCookie("Authorization", token, 3600*24*7, "", "", false, true)
		context.JSON(http.StatusAccepted, gin.H{"message": "User Verified and Logged In"})
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User Not Verified"})
	}
}

func LogoutUser(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	context.SetSameSite(http.SameSiteNoneMode)
	context.SetCookie("Authorization", "", -1, "", "", false, true)
	context.JSON(http.StatusOK, gin.H{"message": "User Logged Out"})
}

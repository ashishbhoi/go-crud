package services

import (
	"github.com/ashishbhoi/go-crud/database"
	"github.com/ashishbhoi/go-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userModel struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func CheckUser(email string) bool {
	var user models.User
	database.DB.Where("email = ?", email).First(&user)
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if CheckUser(user.Email) {
		context.JSON(http.StatusConflict, gin.H{"message": "User Already Exists"})
		return
	}
	database.DB.Create(&user)
	displayUser := userModel{user.FirstName, user.LastName, user.Email}
	context.JSON(201, gin.H{"message": "User Created Successfully", "user": displayUser})
}

func VerifyUser(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var user models.User
	var verifyUser models.User
	err := context.BindJSON(&verifyUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Where("email = ?", verifyUser.Email).First(&user)
	if user.Email == verifyUser.Email && user.Password == verifyUser.Password {
		displayUser := userModel{user.FirstName, user.LastName, user.Email}
		context.JSON(http.StatusAccepted, gin.H{"message": "User Verified", "user": displayUser})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User Not Verified"})
	}
}

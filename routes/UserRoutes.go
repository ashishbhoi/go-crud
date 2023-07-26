package routes

import (
	"github.com/ashishbhoi/go-crud/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/api/users/register", services.CreateUser)
	r.POST("/api/users/login", services.VerifyUser)
}

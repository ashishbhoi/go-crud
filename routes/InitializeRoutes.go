package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	r := gin.Default()
	r.ForwardedByClientIP = true
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})

	UserRoutes(r)
	CategoryRoutes(r)
	TransactionRoutes(r)

	return r
}

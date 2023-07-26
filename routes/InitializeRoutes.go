package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitializeRoutes(port string) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})

	UserRoutes(r)
	CategoryRoutes(r)

	log.Fatal(
		r.Run(port),
	)
}

package routes

import (
	"github.com/ashishbhoi/go-crud/middlewares"
	"github.com/ashishbhoi/go-crud/services"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	r.GET("/api/categories/:categoryId/transactions", middlewares.AuthFilter, services.GetAllTransactions)
	r.GET("/api/categories/:categoryId/transactions/:transactionId", middlewares.AuthFilter, services.GetTransactionById)
	r.POST("/api/categories/:categoryId/transactions", middlewares.AuthFilter, services.CreateTransaction)
	r.PUT("/api/categories/:categoryId/transactions/:transactionId", middlewares.AuthFilter, services.UpdateTransaction)
	r.DELETE("/api/categories/:categoryId/transactions/:transactionId", middlewares.AuthFilter, services.DeleteTransaction)
}

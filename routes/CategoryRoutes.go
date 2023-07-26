package routes

import (
	"github.com/ashishbhoi/go-crud/middlewares"
	"github.com/ashishbhoi/go-crud/services"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	r.GET("/api/categories", middlewares.AuthFilter, services.GetAllCategories)
	r.GET("/api/categories/:categoryId", middlewares.AuthFilter, services.GetCategoryById)
	r.POST("/api/categories", middlewares.AuthFilter, services.CreateCategory)
	r.PUT("/api/categories/:categoryId", middlewares.AuthFilter, services.UpdateCategory)
	r.DELETE("/api/categories/:categoryId", middlewares.AuthFilter, services.DeleteCategory)
}

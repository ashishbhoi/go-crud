package services

import (
	"github.com/ashishbhoi/go-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoryModel struct {
	ID           uint    `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	TotalExpense float64 `json:"total_expense"`
}

func GetAllCategories(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var categories []models.Category
	user, _ := context.Get("user")
	models.DB.Where("user_id = ?", user.(models.PublicUser).ID).Find(&categories)
	var displayCategories []categoryModel
	for _, category := range categories {
		displayCategory := categoryModel{category.ID, category.Title, category.Description, 0.0}
		displayCategories = append(displayCategories, displayCategory)
	}
	context.JSON(http.StatusOK, gin.H{"categories": displayCategories})
}

func GetCategory(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var category models.Category
	categoryId := context.Param("id")
	user, _ := context.Get("user")
	models.DB.Where("id = ? AND user_id = ?", categoryId, user.(models.PublicUser).ID).First(&category)
	displayCategory := categoryModel{category.ID, category.Title, category.Description, 0.0}
	context.JSON(http.StatusOK, gin.H{"category": displayCategory})
}

func CreateCategory(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var category models.Category
	err := context.BindJSON(&category)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, _ := context.Get("user")
	category.UserId = user.(models.PublicUser).ID
	models.DB.Create(&category)
	displayCategory := categoryModel{category.ID, category.Title, category.Description, 0.0}
	context.JSON(http.StatusCreated, gin.H{"message": "Category Created Successfully", "category": displayCategory})
}

func UpdateCategory(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var category models.Category
	categoryId := context.Param("id")
	user, _ := context.Get("user")
	models.DB.Where("id = ? AND user_id = ?", categoryId, user.(models.PublicUser).ID).First(&category)
	var newCategory models.Category
	err := context.BindJSON(&newCategory)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newCategory.Title != "" {
		category.Title = newCategory.Title
	}
	if newCategory.Description != "" {
		category.Description = newCategory.Description
	}
	models.DB.Save(&category)
	context.JSON(http.StatusOK, gin.H{"message": "Category Updated Successfully", "success": true})
}

func DeleteCategory(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var category models.Category
	categoryId := context.Param("id")
	user, _ := context.Get("user")
	models.DB.Where("id = ? AND user_id = ?", categoryId, user.(models.PublicUser).ID).First(&category)
	if category.ID == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Category Not Found", "success": false})
		return
	}
	models.DB.Delete(&category)
	context.JSON(http.StatusOK, gin.H{"message": "Category Deleted Successfully", "success": true})
}

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
	TotalExpense float32 `json:"total_expense"`
}

func GetAllCategories(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var categories []models.Category
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	models.DB.Where("user_id = ?", userId).Find(&categories)
	var displayCategories []categoryModel
	for _, category := range categories {

		var totalExpense float32 = 0.0
		var transactions []models.Transaction
		models.DB.Where("category_id = ? AND user_id = ?", category.ID, userId).Find(&transactions)
		for _, transaction := range transactions {
			totalExpense += transaction.Amount
		}

		displayCategory := categoryModel{category.ID, category.Title, category.Description, totalExpense}
		displayCategories = append(displayCategories, displayCategory)
	}
	context.JSON(http.StatusOK, gin.H{"categories": displayCategories})
}

func GetCategoryById(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var category models.Category
	categoryId := context.Param("categoryId")
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	models.DB.Where("id = ? AND user_id = ?", categoryId, userId).First(&category)

	var totalExpense float32 = 0.0
	var transactions []models.Transaction
	models.DB.Where("category_id = ? AND user_id = ?", category.ID, userId).Find(&transactions)
	for _, transaction := range transactions {
		totalExpense += transaction.Amount
	}

	displayCategory := categoryModel{category.ID, category.Title, category.Description, totalExpense}
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
	categoryId := context.Param("categoryId")
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	models.DB.Where("id = ? AND user_id = ?", categoryId, userId).First(&category)
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
	categoryId := context.Param("categoryId")
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	models.DB.Where("id = ? AND user_id = ?", categoryId, userId).First(&category)
	if category.ID == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Category Not Found", "success": false})
		return
	}
	DeleteAllTransactions(userId, category.ID)
	models.DB.Delete(&category)
	context.JSON(http.StatusOK, gin.H{"message": "Category Deleted Successfully", "success": true})
}

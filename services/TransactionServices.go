package services

import (
	"github.com/ashishbhoi/go-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type transactionModel struct {
	ID              uint    `json:"id"`
	Amount          float32 `json:"amount"`
	Note            string  `json:"note"`
	TransactionDate int64   `json:"transactionDate"`
}

func GetAllTransactions(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var transactions []models.Transaction
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	categoryId := context.Param("categoryId")
	models.DB.Where("category_id = ? AND user_id = ?", categoryId, userId).Find(&transactions)
	var displayTransactions []transactionModel
	for _, transaction := range transactions {
		displayTransaction := transactionModel{transaction.ID, transaction.Amount, transaction.Note,
			transaction.TransactionDate}
		displayTransactions = append(displayTransactions, displayTransaction)
	}
	context.JSON(http.StatusOK, gin.H{"transactions": displayTransactions})
}

func GetTransactionById(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var transaction models.Transaction
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	categoryId := context.Param("categoryId")
	transactionId := context.Param("transactionId")
	models.DB.Where("id = ? AND category_id = ? AND user_id = ?", transactionId, categoryId,
		userId).First(&transaction)
	displayTransaction := transactionModel{transaction.ID, transaction.Amount, transaction.Note,
		transaction.TransactionDate}
	context.JSON(http.StatusOK, gin.H{"transaction": displayTransaction})
}

func CreateTransaction(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var transaction models.Transaction
	err := context.BindJSON(&transaction)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	categoryId := context.Param("categoryId")
	categoryIdU64, _ := strconv.ParseUint(categoryId, 10, 32)
	transaction.CategoryId = uint(categoryIdU64)
	transaction.UserId = userId
	models.DB.Create(&transaction)
	displayTransaction := transactionModel{transaction.ID, transaction.Amount, transaction.Note,
		transaction.TransactionDate}
	context.JSON(http.StatusCreated, gin.H{"message": "Transaction Created Successfully",
		"transaction": displayTransaction})
}

func UpdateTransaction(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var transaction, newTransaction models.Transaction
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	categoryId := context.Param("categoryId")
	transactionId := context.Param("transactionId")
	models.DB.Where("id = ? AND category_id = ? AND user_id = ?", transactionId, categoryId,
		userId).First(&transaction)
	err := context.BindJSON(&newTransaction)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if newTransaction.Amount != 0 {
		transaction.Amount = newTransaction.Amount
	}
	if newTransaction.Note != "" {
		transaction.Note = newTransaction.Note
	}
	if newTransaction.TransactionDate != 0 {
		transaction.TransactionDate = newTransaction.TransactionDate
	}
	models.DB.Save(&transaction)
	context.JSON(http.StatusOK, gin.H{"message": "Transaction Updated Successfully", "success": true})
}

func DeleteTransaction(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var transaction models.Transaction
	user, _ := context.Get("user")
	userId := user.(models.PublicUser).ID
	categoryId := context.Param("categoryId")
	transactionId := context.Param("transactionId")
	models.DB.Where("id = ? AND category_id = ? AND user_id = ?", transactionId, categoryId,
		userId).First(&transaction)
	models.DB.Delete(&transaction)
	context.JSON(http.StatusOK, gin.H{"message": "Transaction Deleted Successfully", "success": true})
}

func DeleteAllTransactions(userId uint, categoryId uint) {
	var transactions []models.Transaction
	models.DB.Where("category_id = ? AND user_id = ?", categoryId, userId).Find(&transactions)
	if len(transactions) == 0 {
		return
	}
	for _, transaction := range transactions {
		models.DB.Delete(&transaction)
	}
}

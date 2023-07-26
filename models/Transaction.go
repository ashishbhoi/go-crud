package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount          float32 `json:"amount"`
	Note            string  `json:"note"`
	TransactionDate int64   `json:"transactionDate"`

	CategoryId uint `json:"category_id"`
	UserId     uint `json:"user_id"`
}

package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`

	UserId       uint          `json:"user_id"`
	Transactions []Transaction `json:"transactions"`
}

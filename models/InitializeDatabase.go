package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func InitializeDatabase() {
	DNS := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" +
		os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE") + " TimeZone=" + os.Getenv("DB_TIMEZONE")

	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	err := DB.AutoMigrate(&User{}, &Category{}, &Transaction{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to migrate database!")
	}
}

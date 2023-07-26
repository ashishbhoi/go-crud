package database

import (
	"fmt"
	"github.com/ashishbhoi/go-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error
var DNS = "host=" + os.Getenv("host") + " user=" + os.Getenv("user") + " password=" +
	os.Getenv("password") + " dbname=" + os.Getenv("db") + " port=" + os.Getenv("port") +
	" sslmode=disable" + " TimeZone=" + os.Getenv("timezone")

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to migrate User!")
	}
	err = DB.AutoMigrate(&models.Category{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to migrate Category!")
	}
	err = DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to migrate Transaction!")
	}
}

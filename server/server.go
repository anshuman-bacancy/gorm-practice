package server

import (
	"fake.com/anshuman/utils"
	"fake.com/anshuman/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Db is global database struct
var Db *gorm.DB
var dbError error

// InitializeDatabase initializes database
func InitializeDatabase() {
	dsn := "anshuman:anshuman32@tcp(localhost:3306)/bookDb?charset=utf8mb4&parseTime=True&loc=Local"
  Db, dbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.CheckError((dbError))

	// Migrations
	Db.AutoMigrate(models.Book{})
}

// InitializeFileServer initializes file server
func InitializeFileServer() {
}

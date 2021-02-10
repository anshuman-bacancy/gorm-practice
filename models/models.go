package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name string `gorm:"column:Name"`
	Author string `gorm:"column:Author"`
	Pages int64 `gorm:"column:Pages"`
	ISBN int64 `gorm:"column:ISBN"`
}


/*
func MakeMigrations() {
	server.Db.AutoMigrate(&Book{})
}
*/

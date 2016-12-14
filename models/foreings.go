package models

import "github.com/jinzhu/gorm"

func CreateForeings(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

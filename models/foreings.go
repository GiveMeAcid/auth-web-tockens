package models

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

func Migrations(db *gorm.DB) {
	//services.DB.DropTableIfExists(&User{}, "users")
	db.AutoMigrate(&User{}).Debug()
	//services.DB.CreateTable(&User{})
}



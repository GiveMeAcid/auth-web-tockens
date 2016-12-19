package services

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB() {
	db, _ := gorm.Open("postgres", "host=localhost user=postgres dbname=auth-web-tokens sslmode=disable password=31780")

	DB = db

	fmt.Println("Connected to the database was succusfully!")
}

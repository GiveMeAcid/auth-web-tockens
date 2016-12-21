package services

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB() error {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=web-tokens sslmode=disable password=31780")

	if err != nil {
		return  err
	}

	db.SingularTable(true)
	fmt.Println("Connected to the database was succusfully!")

	return nil
}

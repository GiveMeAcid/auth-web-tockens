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

	//if err != nil {
	//	fmt.Printf("Database opening error -->%v\n", err)
	//	panic("Database error")
	//}
	//defer db.Close()
	//
	//db.SingularTable(true)

	fmt.Println("Connected to the database was succusfully!")
}

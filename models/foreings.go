package models

import (
	"github.com/satori/go.uuid"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

func Migrations(db *gorm.DB) {
	//services.DB.DropTableIfExists(&User{}, "users")
	db.AutoMigrate(&User{}).Debug()
	//services.DB.CreateTable(&User{})
	//db.Exec("INSERT INTO user VALUES ('7bee9999-230e-47a2-aa4b-351c846a3262', '$2a$10$IRkp9aHaZQJuqgdoMLMLdOfojqMsRp4Infad9aiKrnYybxv1bgTZS')")
	id := uuid.FromStringOrNil("7bee9999-230e-47a2-aa4b-351c846a3262")
	db.Create(&User{
		UUID: id,
		Email: "user21@email.com",
		Password: "$2a$10$IRkp9aHaZQJuqgdoMLMLdOfojqMsRp4Infad9aiKrnYybxv1bgTZS",
		Name: "User 1",
	})
}





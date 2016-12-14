package models

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/auth-web-tokens/services"
)

// fill tables with test data
func Fixtures() {

	test := createUser("test@gmail.com", "1111", USER_ROLE_USER)
	user := createUser("user@gmail.com", "1111", USER_ROLE_USER)
	admin := createUser("admin@gmail.com", "1111", USER_ROLE_ADMIN)

	services.DB.Save(test)
	services.DB.Save(user)
	services.DB.Save(admin)
}

func Migrations() {
	//services.DB.DropTableIfExists(&User{}, "users")
	services.DB.AutoMigrate(&User{}).Debug()
	//services.DB.CreateTable(&User{})
}



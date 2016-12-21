package models

import (
	"github.com/satori/go.uuid"
	"github.com/auth-web-tokens/services"
)


type User struct {
	UUID     uuid.UUID `gorm:"type:uuid;index:idx_user_uuid;not null;column:uuid" json:"uuid"`
	Id       uint `gorm:"primary_key;not null"json:"id"`
	Name     string `gorm:"type:varchar(64);not null"json:"name,omitempty"`
	Email    string `gorm:"type:varchar(64);not null"json:"email,omitempty" form:"email"`
	Password string `gorm:"type:varchar(64);not null"json:"password,omitempty" form:"password"`
	Age      int    `json:"age,omitempty"`
	Role     int `gorm:"not null"`
}

type Users []User

func (user *User) GetById(uuid uuid.UUID) error {
	err := services.DB.Where(&User{UUID: uuid}).First(user).Error
	if err == nil {
		panic(err)
	}
	return err
}

//func (user *UserInfo) Get(email string) error {
//	err := services.DB.Where("e_mail = ?", email).First(user).Error
//	if err == nil {
//		err = user.UserSettings.GetUserSettigs(user.UserSettingsFk)
//		if err == nil {
//			err = user.UserFilters.GetUserFilters(user.UserFiltersFk)
//			if err == nil {
//				err = user.Phone.GetPhone(user.PhoneFk)
//			}
//		}
//	}
//
//	return err
//}





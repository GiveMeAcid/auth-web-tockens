package models

import (
)

const (
	USER_ROLE_USER int = 1
	USER_ROLE_ADMIN int = 2
	USER_AUTH_TOKEN_LIFETIME_SECONDS int = 60 * 60 * 24
)

type User struct {
	UUID     string `json:"uuid" form:"-"`
	Id       uint `gorm:"primary_key;not null"json:"id"`
	Name     string `gorm:"type:varchar(64);not null"json:"name,omitempty"`
	Email    string `gorm:"type:varchar(64);not null"json:"email,omitempty" form:"email"`
	Password string `gorm:"type:varchar(64);not null"json:"password,omitempty" form:"password"`
	Age      int    `json:"age,omitempty"`
	Role     int `gorm:"not null"`
	//AuthToken          string `gorm:"default: null"json:"-"`
	//AuthTokenExpiredAt time.Time `gorm:"default: null"json:"-"`
}

type Users []User


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





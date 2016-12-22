package models

import (
	"github.com/satori/go.uuid"
	"github.com/auth-web-tokens/services"
)

type User struct {
	UUID     uuid.UUID `gorm:"type:uuid;index:idx_user_uuid;not null;column:uuid" json:"uuid"`
	Id       uint `gorm:"primary_key:true;index:idx_user_id;auto_increment:true;column:id" json:"-"`
	Name     string `gorm:"type:varchar(64);not null"json:"name,omitempty"`
	Password string `gorm:"type:varchar(64);not null" validate:"nonzero,max=50" json:"password,omitempty"`
	Email    string `gorm:"type:varchar(255)" validate:"nonzero,max=255,regexp=^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$" json:"email"`
}

type Users []User

func (user *User) GetById(uuid uuid.UUID) error {
	err := services.DB.Where(&User{UUID: uuid}).First(user).Error
	if err == nil {
		panic(err)
	}
	return err
}

func (user *User) Get(email string) error {
	err := services.DB.Where("email = ?", email).First(user).Error
	if err == nil {

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





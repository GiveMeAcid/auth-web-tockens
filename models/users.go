package models

import (
	"time"
	"github.com/rs/xid"
	"github.com/auth-web-tokens/services"
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

func (u *User) SetPassword(plainPassword string) {
	u.Password = services.ToSha1(plainPassword)
}

func (u *User) CheckIsPasswordValid(plainPassword string) bool {
	return u.Password == services.ToSha1(plainPassword)
}

// generates auth token and expired time
func (u *User) GenerateAuthTokenData() {

	var (
		guid xid.ID = xid.New()
		duration time.Duration = time.Duration(USER_AUTH_TOKEN_LIFETIME_SECONDS)
		expiredAt time.Time = time.Now()
	)

	expiredAt = expiredAt.Add(duration) // token will be expire in 24 hours

	u.AuthToken = services.ToSha1(guid.String())
	u.AuthTokenExpiredAt = expiredAt
}

func (u User) IsAuthTokenExpired() bool {
	return u.AuthTokenExpiredAt.After(time.Now())
}

func createUser(email, password string, role int) *User {

	user := &User{
		Email:email,
		Role: role,
	}

	user.SetPassword(password)

	return user
}






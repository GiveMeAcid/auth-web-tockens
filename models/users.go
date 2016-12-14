package models

import (
	//"database/sql/driver"
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
	Id                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Email              string `json:"email,omitempty"`
	Password           string `json:"password,omitempty"`
	Age                int    `json:"age,omitempty"`
	//Gender             Gender `json:"gender,omitempty"`
	Role               int `gorm:"not null"`
	AuthToken          string `gorm:"default: null"json:"-"`
	AuthTokenExpiredAt time.Time `gorm:"default: null"json:"-"`
}

//const (
//	male Gender = "male"
//	female Gender = "female"
//)
//
//func (u *Gender) Scan(value interface{}) error {
//	*u = Gender(value.(string)); return nil
//}
//
//func (u *Gender) Value() (driver.Value, error) {
//	return string(u), nil
//}

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

// fill tables with test data
func Fixtures() {

	test := createUser("test@gmail.com", "1111", USER_ROLE_USER)
	user := createUser("user@gmail.com", "1111", USER_ROLE_USER)
	admin := createUser("admin@gmail.com", "1111", USER_ROLE_ADMIN)

	services.DB.Save(test)
	services.DB.Save(user)
	services.DB.Save(admin)
}

// create tables
func Migrations() {
	services.DB.DropTableIfExists(&User{}, "users")
	services.DB.CreateTable(&User{})
}

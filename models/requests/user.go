package requests

import (
	"errors"
)

type User struct {
	UUID     string `json:"uuid" form:"-"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var (
	UserNotFound = errors.New("User not found")
	EmailUnavailableError = errors.New("Email already used")
	InvalidPassworError = errors.New("Invalid current password")
)
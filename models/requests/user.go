package requests

type User struct {
	UUID     string `json:"uuid" form:"-"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

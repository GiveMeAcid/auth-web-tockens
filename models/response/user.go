package response

import (
	uuid "github.com/satori/go.uuid"
	"github.com/auth-web-tokens/models"
)

type User struct {
	UUID     uuid.UUID `json:"uuid"`
}

type ProfileInfo struct {
	UUID        uuid.UUID     `json:"uuid"`
	EMail       string        `json:"email"`
	Token       string        `json:"token,omitempty"`
}

func NewProfileInfo(user *models.User) *ProfileInfo {
	return &ProfileInfo{
		UUID:        user.UUID,
		EMail:       user.Email,
	}
}
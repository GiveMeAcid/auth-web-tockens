package models

import (
	"time"
)

type Conversation struct {
	ID        uint `gorm:"primary_key;unique"`
	Users     []User
	Messages  []Message
	CreatedAt time.Time `sql:"index"`
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Message struct {
	ID             uint `gorm:"primary_key;unique"`
	ConversationID uint
	Text           string
	CreatedAt      time.Time `sql:"index"`
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}
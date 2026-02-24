package user_entities

import (
	"github.com/go-minstack/sqlite"
)

type User struct {
	sqlite.UuidModel
	Name  string `gorm:"not null"`
	Email string `gorm:"uniqueIndex;not null"`
}

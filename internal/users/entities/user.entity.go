package user_entities

import (
	"github.com/go-minstack/repository"
	"github.com/go-minstack/sqlite"
	"gorm.io/gorm"
)

type User struct {
	sqlite.UuidModel
	Name  string `gorm:"not null"`
	Email string `gorm:"uniqueIndex;not null"`
}

type UserRepository struct {
	*repository.UuidRepository[User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{(*repository.UuidRepository[User])(repository.NewUuidRepository[User](db))}
}

// Domain-specific query — goes here, not in the service
func (r *UserRepository) FindByEmail(email string) (*User, error) {
	return r.FindOne(repository.Where("email = ?", email))
}

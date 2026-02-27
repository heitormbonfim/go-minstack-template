package user_repositories

import (
	user_entities "go-minstack-task/internal/users/entities"

	"github.com/go-minstack/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	*repository.Repository[user_entities.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		repository.NewRepository[user_entities.User](db),
	}
}

func (r *UserRepository) FindByEmail(email string) (*user_entities.User, error) {
	return r.FindOne(repository.Where("email = ?", email))
}

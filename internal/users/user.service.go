package users

import (
	"go-minstack/internal/users/dto"
	user_entities "go-minstack/internal/users/entities"
	"go-minstack/internal/users/repositories"
	"log/slog"

	"github.com/go-minstack/repository"
	"github.com/google/uuid"
)

type UserService struct {
	log   *slog.Logger
	users *repositories.UserRepository
}

func NewUserService(log *slog.Logger, users *repositories.UserRepository) *UserService {
	return &UserService{
		log:   log,
		users: users,
	}
}

func (s *UserService) Create(input dto.CreateUserDto) (*dto.UserDto, error) {
	user := &user_entities.User{
		Name:  input.Name,
		Email: input.Email,
	}
	if err := s.users.Save(user); err != nil {
		return nil, err
	}
	result := dto.NewUserDto(user)
	s.log.Info("user created", "name", result.Name)
	return &result, nil
}

func (s *UserService) List() ([]dto.UserDto, error) {
	users, err := s.users.FindAll(repository.Order("name"))
	if err != nil {
		return nil, err
	}
	result := make([]dto.UserDto, len(users))
	for i, u := range users {
		result[i] = dto.NewUserDto(&u)
	}
	return result, nil
}

func (s *UserService) FindUserByID(id uuid.UUID) (*dto.UserDto, error) {
	user, err := s.users.FindByID(id)
	if err != nil {
		return nil, err
	}
	result := dto.NewUserDto(user)
	return &result, nil
}

func (s *UserService) Update(id uuid.UUID, input dto.UpdateUserDto) (*dto.UserDto, error) {
	user, err := s.users.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name

	if err := s.users.Update(user); err != nil {
		return nil, err
	}
	result := dto.NewUserDto(user)
	s.log.Info("user updated", "name", result.Name)
	return &result, nil
}

func (s *UserService) Delete(id uuid.UUID) error {
	if err := s.users.DeleteByID(id); err != nil {
		return err
	}
	return nil
}

package users

import (
	"go-minstack/internal/users/dto"
	user_entities "go-minstack/internal/users/entities"

	"github.com/go-minstack/repository"
	"github.com/google/uuid"
)

type UserService struct {
	users *user_entities.UserRepository
}

func NewUserService(users *user_entities.UserRepository) *UserService {
	return &UserService{users: users}
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
	return &result, nil
}

func (s *UserService) List() ([]dto.UserDto, error) {
	users, err := s.users.FindAll(repository.Order("name"))
	if err != nil {
		return nil, err
	}
	dtos := make([]dto.UserDto, len(users))
	for i, u := range users {
		dtos[i] = dto.NewUserDto(&u)
	}
	return dtos, nil
}

func (s *UserService) Update(id string, input dto.UpdateUserDto) (*dto.UserDto, error) {
	user, err := s.users.FindByStringID(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name

	if err := s.users.Update(user); err != nil {
		return nil, err
	}
	result := dto.NewUserDto(user)
	return &result, nil
}

func (s *UserService) Delete(id uuid.UUID) error {
	if err := s.users.DeleteByID(id); err != nil {
		return err
	}
	return nil
}

package users

import (
	user_dto "go-minstack-task/internal/users/dtos"
	user_entities "go-minstack-task/internal/users/entities"
	user_repositories "go-minstack-task/internal/users/repositories"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	users *user_repositories.UserRepository
	log   *slog.Logger
}

func NewUserService(users *user_repositories.UserRepository, log *slog.Logger) *UserService {
	return &UserService{users: users, log: log}
}

func (s *UserService) Register(input user_dto.RegisterDto) (*user_dto.UserDto, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &user_entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
	}
	if err := s.users.Create(user); err != nil {
		s.log.Error("failed to create user", "error", err)
		return nil, err
	}
	s.log.Info("user registered", "user_id", user.ID)
	result := user_dto.NewUserDto(user)
	return &result, nil
}

func (s *UserService) Me(id uint) (*user_dto.UserDto, error) {
	user, err := s.users.FindByID(id)
	if err != nil {
		s.log.Error("user not found", "user_id", id)
		return nil, err
	}
	result := user_dto.NewUserDto(user)
	return &result, nil
}

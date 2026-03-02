package authn

import (
	"errors"
	"fmt"
	user_dto "go-minstack-task/internal/users/dtos"
	user_repositories "go-minstack-task/internal/users/repositories"
	"log/slog"
	"time"

	"github.com/go-minstack/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthnService struct {
	users *user_repositories.UserRepository
	jwt   *auth.JwtService
	log   *slog.Logger
}

func NewAuthnService(users *user_repositories.UserRepository, jwt *auth.JwtService, log *slog.Logger) *AuthnService {
	return &AuthnService{users: users, jwt: jwt, log: log}
}

func (s *AuthnService) Login(input user_dto.LoginDto) (string, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		s.log.Warn("login failed: user not found")
		return "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		s.log.Warn("login failed: wrong password", "user_id", user.ID)
		return "", errors.New("invalid credentials")
	}
	token, err := s.jwt.Sign(auth.Claims{
		Subject: fmt.Sprintf("%d", user.ID),
		Name:    user.Name,
	}, 24*time.Hour)
	if err != nil {
		return "", err
	}
	s.log.Info("user authenticated", "user_id", user.ID)
	return token, nil
}

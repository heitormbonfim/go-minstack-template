package authn

import (
	"errors"
	"fmt"
	user_dto "go-minstack-task/internal/users/dtos"
	user_repositories "go-minstack-task/internal/users/repositories"
	"time"

	"github.com/go-minstack/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthnService struct {
	users *user_repositories.UserRepository
	jwt   *auth.JwtService
}

const invalidCredentials = "invalid credentials"

func NewAuthnService(users *user_repositories.UserRepository, jwt *auth.JwtService) *AuthnService {
	return &AuthnService{users: users, jwt: jwt}
}

func (s *AuthnService) Login(input user_dto.LoginDto) (string, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New(invalidCredentials)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New(invalidCredentials)
	}
	return s.jwt.Sign(auth.Claims{
		Subject: fmt.Sprintf("%d", user.ID),
		Name:    user.Name,
	}, 24*time.Hour)
}

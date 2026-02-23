// dto/user.dto.go
package dto

import user_entities "go-minstack/internal/users/entities"

type UserDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserDto(u *user_entities.User) UserDto {
	return UserDto{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: u.Email,
	}
}

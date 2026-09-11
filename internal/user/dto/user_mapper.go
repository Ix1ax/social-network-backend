package dto

import (
	"github.com/ix1ax/social-network-backend/internal/user/entity"
)

func ToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToAuthResponse(accessToken string, user *entity.User) *AuthResponse {
	return &AuthResponse{
		AccessToken: accessToken,
		User:        ToUserResponse(user),
	}
}

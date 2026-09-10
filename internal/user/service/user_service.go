package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ix1ax/social-network-backend/internal/user/dto"
	"github.com/ix1ax/social-network-backend/internal/user/entity"
	"github.com/ix1ax/social-network-backend/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists = errors.New("user with this email already exists")
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.UserResponse, error) {

	_, err := s.userRepo.GetByEmail(ctx, req.Email)

	if err == nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("failed to hash password %w", err)
	}

	newUser := entity.User{
		Name:         req.Name,
		Surname:      req.Surname,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepo.Create(ctx, &newUser)

	if err != nil {
		return nil, fmt.Errorf("failed to created user %w", err)
	}

	return &dto.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Surname:   newUser.Surname,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}, nil

}

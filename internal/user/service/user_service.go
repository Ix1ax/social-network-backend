package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/ix1ax/social-network-backend/internal/common/token"
	"github.com/ix1ax/social-network-backend/internal/user/dto"
	"github.com/ix1ax/social-network-backend/internal/user/entity"
	"github.com/ix1ax/social-network-backend/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExists  = errors.New("user with this email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not founded")
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.UserResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error)
}

type userService struct {
	userRepo     repository.UserRepository
	tokenManager token.TokenManager
}

func NewUserService(userRepo repository.UserRepository, tokenManager token.TokenManager) UserService {
	return &userService{
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
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

	return dto.ToUserResponse(&newUser), nil

}

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {

	user, err := s.userRepo.GetByEmail(ctx, req.Email)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	jwtToken, err := s.tokenManager.GenerateToken(user.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return dto.ToAuthResponse(jwtToken, user), nil

}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error) {

	user, err := s.userRepo.GetByID(ctx, id)

	if err != nil {
		return nil, ErrUserNotFound
	}

	return dto.ToUserResponse(user), nil
}

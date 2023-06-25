package services

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	UpdateUserByUsername(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(ctx, username)
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	existingUser, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		fmt.Printf("failed to get user by email: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("email already registered")
	}

	return s.userRepo.CreateUser(ctx, user)

	// existingUser, err := s.userRepo.GetUserByID(ctx, user.ID)
	// if err != nil {
	// 	return fmt.Errorf("failed to get user by ID: %v", err)
	// }
	// if existingUser != nil {
	// 	return fmt.Errorf("email already registered")
	// }

	// return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *userService) UpdateUserByUsername(ctx context.Context, user *models.User) error {
	
	return s.userRepo.UpdateUserByUsername(ctx, user.Username, user)

}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.DeleteUser(ctx, id)
}

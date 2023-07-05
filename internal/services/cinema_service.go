package services

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
)

type CinemaService interface {
	GetAllCinemas(ctx context.Context) ([]models.Cinema, error)
	GetCinemaByID(ctx context.Context, id string) (*models.Cinema, error)
	CreateCinema(ctx context.Context, cinema *models.Cinema) error
	UpdateCinema(ctx context.Context, cinema *models.Cinema) error
	DeleteCinema(ctx context.Context, id string) error
}

type cinemaService struct {
	cinemaRepo repositories.CinemaRepository
}

func NewCinemaService(cinemaRepo repositories.CinemaRepository) CinemaService {
	return &cinemaService{
		cinemaRepo: cinemaRepo,
	}
}

func (s *cinemaService) GetAllCinemas(ctx context.Context) ([]models.Cinema, error) {
	return s.cinemaRepo.GetAllCinemas(ctx)
}

func (s *cinemaService) GetCinemaByID(ctx context.Context, id string) (*models.Cinema, error) {
	return s.cinemaRepo.GetCinemaByID(ctx, id)
}

func (s *cinemaService) CreateCinema(ctx context.Context, cinema *models.Cinema) error {
	existingCinema, err := s.cinemaRepo.GetCinemaByID(ctx, cinema.ID.Hex())
	if err != nil {
		return fmt.Errorf("failed to get cinema by id: %v", err)
	}
	if existingCinema != nil {
		return fmt.Errorf("cinema already registered")
	}

	return s.cinemaRepo.CreateCinema(ctx, cinema)
}

func (s *cinemaService) UpdateCinema(ctx context.Context, cinema *models.Cinema) error {
	return s.cinemaRepo.UpdateCinema(ctx, cinema)
}

func (s *cinemaService) DeleteCinema(ctx context.Context, id string) error {
	return s.cinemaRepo.DeleteCinema(ctx, id)
}

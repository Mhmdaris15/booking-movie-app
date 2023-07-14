package services

import (
	"context"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
)

type ShowtimeService interface {
	GetAllShowtimes(ctx context.Context) ([]models.Showtime, error)
	GetShowtimeByID(ctx context.Context, id string) (*models.Showtime, error)
	CreateShowtime(ctx context.Context, showtime *models.Showtime) error
	UpdateShowtime(ctx context.Context, showtime *models.Showtime) error
	DeleteShowtime(ctx context.Context, id string) error
}

type showtimeService struct {
	showtimeRepo repositories.ShowtimeRepository
}

func NewShowtimeService(showtimeRepo repositories.ShowtimeRepository) ShowtimeService {
	return &showtimeService{
		showtimeRepo: showtimeRepo,
	}
}

func (s *showtimeService) GetAllShowtimes(ctx context.Context) ([]models.Showtime, error) {
	return s.showtimeRepo.GetAllShowtimes(ctx)
}

func (s *showtimeService) GetShowtimeByID(ctx context.Context, id string) (*models.Showtime, error) {
	return s.showtimeRepo.GetShowtimeByID(ctx, id)
}

func (s *showtimeService) CreateShowtime(ctx context.Context, showtime *models.Showtime) error {
	// Perform any necessary business logic or validation before calling the repository function
	// For example, check if a showtime with the same ID already exists

	// existingShowtime, err := s.showtimeRepo.GetShowtimeByID(ctx, showtime.ID.Hex())
	// if err == nil {
	// 	return fmt.Errorf("failed to check existing showtime: %v", err)
	// }
	// if existingShowtime == nil {
	// 	return fmt.Errorf("showtime with ID '%s' already exists", showtime.ID)
	// }

	return s.showtimeRepo.CreateShowtime(ctx, showtime)
}

func (s *showtimeService) UpdateShowtime(ctx context.Context, showtime *models.Showtime) error {
	return s.showtimeRepo.UpdateShowtime(ctx, showtime)
}

func (s *showtimeService) DeleteShowtime(ctx context.Context, id string) error {
	return s.showtimeRepo.DeleteShowtime(ctx, id)
}

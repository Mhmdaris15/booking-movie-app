package repositories

import (
	"context"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
)

type ShowtimeRepository interface {
	GetAllShowtimes(ctx context.Context) ([]models.Showtime, error)
	GetShowtimeByID(ctx context.Context, id string) (*models.Showtime, error)
	CreateShowtime(ctx context.Context, showtime *models.Showtime) error
	UpdateShowtime(ctx context.Context, showtime *models.Showtime) error
	DeleteShowtime(ctx context.Context, id string) error
}
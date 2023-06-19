package repositories

import (
	"context"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
)

type CinemaRepository interface {
	GetAllCinemas(ctx context.Context) ([]models.Cinema, error)
	GetCinemaByID(ctx context.Context, id string) (*models.Cinema, error)
	CreateCinema(ctx context.Context, cinema *models.Cinema) error
	UpdateCinema(ctx context.Context, cinema *models.Cinema) error
	DeleteCinema(ctx context.Context, id string) error
}
package repositories

import (
	"context"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
)

type SeatRepository interface {
	GetAllSeats(ctx context.Context) ([]models.Seat, error)
	GetSeatByID(ctx context.Context, id string) (*models.Seat, error)
	CreateSeat(ctx context.Context, seat *models.Seat) error
	UpdateSeat(ctx context.Context, seat *models.Seat) error
	DeleteSeat(ctx context.Context, id string) error
}
package services

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
)

type SeatService interface {
	GetAllSeats(ctx context.Context) ([]models.Seat, error)
	GetSeatByID(ctx context.Context, id string) (*models.Seat, error)
	CreateSeat(ctx context.Context, seat *models.Seat) error
	UpdateSeat(ctx context.Context, seat *models.Seat) error
	DeleteSeat(ctx context.Context, id string) error
}

type seatService struct {
	seatRepo repositories.SeatRepository
}

func NewSeatService(seatRepo repositories.SeatRepository) SeatService {
	return &seatService{
		seatRepo: seatRepo,
	}
}

func (s *seatService) GetAllSeats(ctx context.Context) ([]models.Seat, error) {
	return s.seatRepo.GetAllSeats(ctx)
}

func (s *seatService) GetSeatByID(ctx context.Context, id string) (*models.Seat, error) {
	return s.seatRepo.GetSeatByID(ctx, id)
}

func (s *seatService) CreateSeat(ctx context.Context, seat *models.Seat) error {
	existingSeat, err := s.seatRepo.GetSeatByID(ctx, seat.ID)
	if err != nil {
		return fmt.Errorf("failed to get seat by id: %v", err)
	}
	if existingSeat != nil {
		return fmt.Errorf("seat already registered")
	}

	return s.seatRepo.CreateSeat(ctx, seat)
}

func (s *seatService) UpdateSeat(ctx context.Context, seat *models.Seat) error {
	return s.seatRepo.UpdateSeat(ctx, seat)
}

func (s *seatService) DeleteSeat(ctx context.Context, id string) error {
	return s.seatRepo.DeleteSeat(ctx, id)
}
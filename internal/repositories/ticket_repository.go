package repositories

import (
	"context"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
)

type TicketRepository interface {
	GetAllTickets(ctx context.Context) ([]models.Ticket, error)
	GetTicketByID(ctx context.Context, id string) (*models.Ticket, error)
	CreateTicket(ctx context.Context, ticket *models.Ticket) error
	UpdateTicket(ctx context.Context, ticket *models.Ticket) error
	DeleteTicket(ctx context.Context, id string) error
}
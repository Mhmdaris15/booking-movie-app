package services

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
)

type TicketService interface {
	GetAllTickets(ctx context.Context) ([]models.Ticket, error)
	GetTicketByID(ctx context.Context, id string) (*models.Ticket, error)
	CreateTicket(ctx context.Context, ticket *models.Ticket) error
	UpdateTicket(ctx context.Context, ticket *models.Ticket) error
	DeleteTicket(ctx context.Context, id string) error
}

type ticketService struct {
	ticketRepo repositories.TicketRepository
}

func NewTicketService(ticketRepo repositories.TicketRepository) TicketService {
	return &ticketService{
		ticketRepo: ticketRepo,
	}
}

func (s *ticketService) GetAllTickets(ctx context.Context) ([]models.Ticket, error) {
	return s.ticketRepo.GetAllTickets(ctx)
}

func (s *ticketService) GetTicketByID(ctx context.Context, id string) (*models.Ticket, error) {
	return s.ticketRepo.GetTicketByID(ctx, id)
}

func (s *ticketService) CreateTicket(ctx context.Context, ticket *models.Ticket) error {
	existingTicket, err := s.ticketRepo.GetTicketByID(ctx, ticket.ID.Hex())
	if err != nil {
		return fmt.Errorf("failed to get ticket by ID: %v", err)
	}
	if existingTicket != nil {
		return fmt.Errorf("ticket already exists")
	}

	return s.ticketRepo.CreateTicket(ctx, ticket)
}

func (s *ticketService) UpdateTicket(ctx context.Context, ticket *models.Ticket) error {
	return s.ticketRepo.UpdateTicket(ctx, ticket)
}

func (s *ticketService) DeleteTicket(ctx context.Context, id string) error {
	return s.ticketRepo.DeleteTicket(ctx, id)
}
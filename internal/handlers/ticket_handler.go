package handlers

import (
	"net/http"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/gin-gonic/gin"
)

type TicketHandler interface {
	GetAllTickets(ctx *gin.Context)
	GetAllTicketsByUserID(ctx *gin.Context)
	GetAllTicketsByUsername(ctx *gin.Context)
	GetTicketByID(ctx *gin.Context)
	CreateTicket(ctx *gin.Context)
	UpdateTicket(ctx *gin.Context)
	DeleteTicket(ctx *gin.Context)
}

type ticketHandler struct {
	ticketService services.TicketService
}

func NewTicketHandler(ticketService services.TicketService) TicketHandler {
	return &ticketHandler{
		ticketService: ticketService,
	}
}

func (h *ticketHandler) GetAllTickets(ctx *gin.Context) {
	tickets, err := h.ticketService.GetAllTickets(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}

func (h *ticketHandler) GetTicketByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ticket, err := h.ticketService.GetTicketByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ticket)
}

func (h *ticketHandler) GetAllTicketsByUserID(ctx *gin.Context) {
	id := ctx.Param("id")
	tickets, err := h.ticketService.GetAllTicketsByUserID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (h *ticketHandler) GetAllTicketsByUsername(ctx *gin.Context) {
	username := ctx.Param("id")
	tickets, err := h.ticketService.GetAllTicketsByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (h *ticketHandler) CreateTicket(ctx *gin.Context) {
	var ticket *models.Ticket
	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTicket, err := h.ticketService.CreateTicket(ctx, ticket)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newTicket)
}

func (h *ticketHandler) UpdateTicket(ctx *gin.Context) {
	var ticket *models.Ticket
	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.ticketService.UpdateTicket(ctx, ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ticket)
}

func (h *ticketHandler) DeleteTicket(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.ticketService.DeleteTicket(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Ticket has been deleted"})
}

// Path: internal\handlers\cinema_handler.go

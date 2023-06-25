package handlers

import (
	"net/http"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/gin-gonic/gin"
)

type SeatHandler interface {
	GetAllSeats(ctx *gin.Context)
	GetSeatByID(ctx *gin.Context)
	CreateSeat(ctx *gin.Context)
	UpdateSeat(ctx *gin.Context)
	DeleteSeat(ctx *gin.Context)
}

type seatHandler struct {
	seatService services.SeatService
}

func NewSeatHandler(seatService services.SeatService) SeatHandler {
	return &seatHandler{
		seatService: seatService,
	}
}

func (h *seatHandler) GetAllSeats(ctx *gin.Context) {
	seats, err := h.seatService.GetAllSeats(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, seats)
}

func (h *seatHandler) GetSeatByID(ctx *gin.Context) {
	id := ctx.Param("id")
	seat, err := h.seatService.GetSeatByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, seat)
}

func (h *seatHandler) CreateSeat(ctx *gin.Context) {
	var seat *models.Seat
	if err := ctx.ShouldBindJSON(&seat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.seatService.CreateSeat(ctx, seat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, seat)
}

func (h *seatHandler) UpdateSeat(ctx *gin.Context) {
	var seat *models.Seat
	if err := ctx.ShouldBindJSON(&seat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.seatService.UpdateSeat(ctx, seat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, seat)
}

func (h *seatHandler) DeleteSeat(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.seatService.DeleteSeat(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Seat deleted successfully"})
}

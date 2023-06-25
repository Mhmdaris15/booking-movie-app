package handlers

import (
	"net/http"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/gin-gonic/gin"
)

type ShowtimeHandler interface {
	GetAllShowtimes(ctx *gin.Context)
	GetShowtimeByID(ctx *gin.Context)
	CreateShowtime(ctx *gin.Context)
	UpdateShowtime(ctx *gin.Context)
	DeleteShowtime(ctx *gin.Context)
}

type showtimeHandler struct {
	showtimeService services.ShowtimeService
}

func NewShowtimeHandler(showtimeService services.ShowtimeService) ShowtimeHandler {
	return &showtimeHandler{
		showtimeService: showtimeService,
	}
}

func (h *showtimeHandler) GetAllShowtimes(ctx *gin.Context) {
	showtimes, err := h.showtimeService.GetAllShowtimes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, showtimes)
}

func (h *showtimeHandler) GetShowtimeByID(ctx *gin.Context) {
	id := ctx.Param("id")
	showtime, err := h.showtimeService.GetShowtimeByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, showtime)
}

func (h *showtimeHandler) CreateShowtime(ctx *gin.Context) {
	var showtime *models.Showtime
	if err := ctx.ShouldBindJSON(&showtime); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.showtimeService.CreateShowtime(ctx, showtime); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, showtime)
}

func (h *showtimeHandler) UpdateShowtime(ctx *gin.Context) {
	var showtime *models.Showtime
	if err := ctx.ShouldBindJSON(&showtime); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.showtimeService.UpdateShowtime(ctx, showtime); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, showtime)
}

func (h *showtimeHandler) DeleteShowtime(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.showtimeService.DeleteShowtime(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "showtime deleted successfully"})
}

// Path: internal\handlers\user_handler.go
package handlers

import (
	"net/http"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/gin-gonic/gin"
)

type CinemaHandler interface {
	GetAllCinemas(ctx *gin.Context)
	GetCinemaByID(ctx *gin.Context)
	CreateCinema(ctx *gin.Context)
	UpdateCinema(ctx *gin.Context)
	DeleteCinema(ctx *gin.Context)
}

type cinemaHandler struct {
	cinemaService services.CinemaService
}

func NewCinemaHandler(cinemaService services.CinemaService) CinemaHandler {
	return &cinemaHandler{
		cinemaService: cinemaService,
	}
}

func (h *cinemaHandler) GetAllCinemas(ctx *gin.Context) {
	cinemas, err := h.cinemaService.GetAllCinemas(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cinemas)
}

func (h *cinemaHandler) GetCinemaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	cinema, err := h.cinemaService.GetCinemaByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cinema)
}

func (h *cinemaHandler) CreateCinema(ctx *gin.Context) {
	var cinema *models.Cinema
	if err := ctx.ShouldBindJSON(&cinema); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.cinemaService.CreateCinema(ctx, cinema); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, cinema)
}

func (h *cinemaHandler) UpdateCinema(ctx *gin.Context) {
	var cinema *models.Cinema
	if err := ctx.ShouldBindJSON(&cinema); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.cinemaService.UpdateCinema(ctx, cinema); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cinema)
}

func (h *cinemaHandler) DeleteCinema(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.cinemaService.DeleteCinema(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cinema deleted successfully"})
}

// Path: internal\handlers\movie_handler.go
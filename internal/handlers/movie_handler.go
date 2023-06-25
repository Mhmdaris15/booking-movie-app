package handlers

import (
	"net/http"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	GetAllMovies(ctx *gin.Context)
	GetMovieByID(ctx *gin.Context)
	GetMovieByTitle(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
}

type movieHandler struct {
	movieService services.MovieService
}

func NewMovieHandler(movieService services.MovieService) MovieHandler {
	return &movieHandler{
		movieService: movieService,
	}
}

func (h *movieHandler) GetAllMovies(ctx *gin.Context) {
	movies, err := h.movieService.GetAllMovies(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (h *movieHandler) GetMovieByID(ctx *gin.Context) {
	id := ctx.Param("id")
	movie, err := h.movieService.GetMovieByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (h *movieHandler) GetMovieByTitle(ctx *gin.Context) {
	title := ctx.Param("title")
	movie, err := h.movieService.GetMovieByTitle(ctx, title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (h *movieHandler) CreateMovie(ctx *gin.Context) {
	var movie *models.Movie
	err := ctx.BindJSON(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.movieService.CreateMovie(ctx, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, movie)
}

func (h *movieHandler) UpdateMovie(ctx *gin.Context) {
	var movie *models.Movie
	err := ctx.BindJSON(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.movieService.UpdateMovie(ctx, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (h *movieHandler) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.movieService.DeleteMovie(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
}
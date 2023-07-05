package handlers

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/response"
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

	responses := []response.Response{}

	for _, movie := range movies {
		imageBytes, err := ioutil.ReadFile(movie.PosterURL)
		if err != nil {
			// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Read File" + err.Error()})
			// return
			imageBytes, _ = ioutil.ReadFile("/uploads/polars.png")
		}
		responses = append(responses, response.Response{
			Data:  movie,
			Image: base64.StdEncoding.EncodeToString(imageBytes),
		})
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Content-Disposition", "attachment; filename=movie.json")

	ctx.JSON(http.StatusOK, responses)
}


func (h *movieHandler) GetMovieByID(ctx *gin.Context) {
	id := ctx.Param("id")
	movie, err := h.movieService.GetMovieByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res response.Response

	imageBytes, err := ioutil.ReadFile(movie.PosterURL)
	if err != nil {
		res = response.Response{
			Data: *movie,
			Image: "",
		}
	} else {
		res = response.Response{
			Data: *movie,
			Image: base64.StdEncoding.EncodeToString(imageBytes),
		}
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Content-Disposition", "attachment; filename=movie.json")


	ctx.JSON(http.StatusOK, res)
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

	file, err := ctx.FormFile("image_file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Get Image File Error: " + err.Error()})
		return
	}

	// Create a new file in the uploads directory
	dstPath := "uploads/" + file.Filename
	dst, err := os.Create(dstPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Create Destination Path" + err.Error()})
		return
	}
	defer dst.Close()

	// Open the file from the source
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// Copy the file to the destination
	_, err = io.Copy(dst, src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove file_image request 
	
	if err = ctx.Request.MultipartForm.RemoveAll(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	num, err := strconv.ParseFloat(ctx.PostForm("ticket_price"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ticket Price Parsing Error: " + err.Error()})
		return
	}

	movie = &models.Movie{
		Title:       ctx.PostForm("title"),
		Description: ctx.PostForm("description"),
		AgeRating:   ctx.PostForm("age_rating"),
		TicketPrice: num,
		PosterURL:   dstPath,
	}

	err = h.movieService.CreateMovie(ctx, movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "movie created successfully", "movie": movie})
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
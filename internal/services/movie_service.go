package services

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
)

type MovieService interface {
	GetAllMovies(ctx context.Context) ([]models.Movie, error)
	GetMovieByID(ctx context.Context, id string) (*models.Movie, error)
	GetMovieByTitle(ctx context.Context, title string) (*models.Movie, error)
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	DeleteMovie(ctx context.Context, id string) error
}

type movieService struct {
	movieRepo repositories.MovieRepository
}

func NewMovieService(movieRepo repositories.MovieRepository) MovieService {
	return &movieService{
		movieRepo: movieRepo,
	}
} 

func (s *movieService) GetAllMovies(ctx context.Context) ([]models.Movie, error) {
	return s.movieRepo.GetAllMovies(ctx)
}

func (s *movieService) GetMovieByID(ctx context.Context, id string) (*models.Movie, error) {
	return s.movieRepo.GetMovieByID(ctx, id)
}

func (s *movieService) GetMovieByTitle(ctx context.Context, title string) (*models.Movie, error) {
	return s.movieRepo.GetMovieByTitle(ctx, title)
}

func (s *movieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	// Perform any necessary business logic or validation before calling the repository function
	// For example, check if a movie with the same title already exists

	existingMovie, err := s.movieRepo.GetMovieByTitle(ctx, movie.Title)
	if err != nil {
		fmt.Printf("failed to check existing movie: %v", err)
	}
	if existingMovie != nil {
		return fmt.Errorf("movie with title '%s' already exists", movie.Title)
	}

	return s.movieRepo.CreateMovie(ctx, movie)
}

func (s *movieService) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	return s.movieRepo.UpdateMovie(ctx, movie)
}

func (s *movieService) DeleteMovie(ctx context.Context, id string) error {
	return s.movieRepo.DeleteMovie(ctx, id)
}

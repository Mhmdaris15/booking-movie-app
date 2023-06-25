package repositories

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository interface {
	GetAllMovies(ctx context.Context) ([]models.Movie, error)
	GetMovieByID(ctx context.Context, id string) (*models.Movie, error)
	GetMovieByTitle(ctx context.Context, title string) (*models.Movie, error)
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	DeleteMovie(ctx context.Context, id string) error
}

type movieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(collection *mongo.Collection) MovieRepository {
	return &movieRepository{
		collection: collection,
	}
}

func (r *movieRepository) GetAllMovies(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve movies: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var movie models.Movie
		if err := cursor.Decode(&movie); err != nil {
			return nil, fmt.Errorf("failed to decode movie: %v", err)
		}
		movies = append(movies, movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return movies, nil
}

func (r *movieRepository) GetMovieByID(ctx context.Context, id string) (*models.Movie, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid movie ID: %v", err)
	}

	filter := bson.M{"_id": objectID}

	var movie models.Movie
	err = r.collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("movie not found")
		}
		return nil, fmt.Errorf("failed to get movie: %v", err)
	}

	return &movie, nil
}

func (r *movieRepository) GetMovieByTitle(ctx context.Context, title string) (*models.Movie, error) {
	filter := bson.M{"title": title}

	var movie models.Movie
	err := r.collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("movie not found")
		}
		return nil, fmt.Errorf("failed to get movie: %v", err)
	}

	return &movie, nil
}

func (r *movieRepository) CreateMovie(ctx context.Context, movie *models.Movie) error {
	_, err := r.collection.InsertOne(ctx, movie)
	if err != nil {
		return fmt.Errorf("failed to create movie: %v", err)
	}

	return nil
}

func (r *movieRepository) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	filter := bson.M{"_id": movie.ID}
	update := bson.M{"$set": movie}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update movie: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("movie not found")
	}

	return nil
}

func (r *movieRepository) DeleteMovie(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid movie ID: %v", err)
	}

	filter := bson.M{"_id": objectID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("movie not found")
	}

	return nil
}
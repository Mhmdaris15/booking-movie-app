package repositories

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CinemaRepository interface {
	GetAllCinemas(ctx context.Context) ([]models.Cinema, error)
	GetCinemaByID(ctx context.Context, id string) (*models.Cinema, error)
	CreateCinema(ctx context.Context, cinema *models.Cinema) error
	UpdateCinema(ctx context.Context, cinema *models.Cinema) error
	DeleteCinema(ctx context.Context, id string) error
}

type cinemaRepository struct {
	collection *mongo.Collection
}

func NewCinemaRepository(collection *mongo.Collection) CinemaRepository {
	return &cinemaRepository{
		collection: collection,
	}
}

func (r *cinemaRepository) GetAllCinemas(ctx context.Context) ([]models.Cinema, error) {
	var cinemas []models.Cinema

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cinemas: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var cinema models.Cinema
		if err := cursor.Decode(&cinema); err != nil {
			return nil, fmt.Errorf("failed to decode cinema: %v", err)
		}
		cinemas = append(cinemas, cinema)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return cinemas, nil
}

func (r *cinemaRepository) GetCinemaByID(ctx context.Context, id string) (*models.Cinema, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid cinema ID: %v", err)
	}
	filter := bson.M{"_id": objectID}

	var cinema models.Cinema
	if err := r.collection.FindOne(ctx, filter).Decode(&cinema); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("cinema not found: %v", err)
		}
		return nil, fmt.Errorf("failed to retrieve cinema: %v", err)
	}

	return &cinema, nil
}

func (r *cinemaRepository) CreateCinema(ctx context.Context, cinema *models.Cinema) error {
	_, err := r.collection.InsertOne(ctx, cinema)
	if err != nil {
		return fmt.Errorf("failed to insert cinema: %v", err)
	}

	return nil
}

func (r *cinemaRepository) UpdateCinema(ctx context.Context, cinema *models.Cinema) error {
	filter := bson.M{"_id": cinema.ID}
	update := bson.M{"$set": cinema}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update cinema: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("cinema not found: %v", err)
	}

	return nil
}

func (r *cinemaRepository) DeleteCinema(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid cinema ID: %v", err)
	}
	filter := bson.M{"_id": objectID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete cinema: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("cinema not found: %v", err)
	}

	return nil
}
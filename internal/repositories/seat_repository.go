package repositories

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SeatRepository interface {
	GetAllSeats(ctx context.Context) ([]models.Seat, error)
	GetSeatByID(ctx context.Context, id string) (*models.Seat, error)
	CreateSeat(ctx context.Context, seat *models.Seat) error
	UpdateSeat(ctx context.Context, seat *models.Seat) error
	DeleteSeat(ctx context.Context, id string) error
}

type seatRepository struct {
	collection *mongo.Collection
}

func NewSeatRepository(collection *mongo.Collection) SeatRepository {
	return &seatRepository{
		collection: collection,
	}
}

func (r *seatRepository) GetAllSeats(ctx context.Context) ([]models.Seat, error) {
	var seats []models.Seat

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve seats: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var seat models.Seat
		if err := cursor.Decode(&seat); err != nil {
			return nil, fmt.Errorf("failed to decode seat: %v", err)
		}
		seats = append(seats, seat)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return seats, nil
}

func (r *seatRepository) GetSeatByID(ctx context.Context, id string) (*models.Seat, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid seat ID: %v", err)
	}

	var seat models.Seat
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&seat); err != nil {
		return nil, fmt.Errorf("failed to retrieve seat: %v", err)
	}

	return &seat, nil
}

func (r *seatRepository) CreateSeat(ctx context.Context, seat *models.Seat) error {
	_, err := r.collection.InsertOne(ctx, seat)
	if err != nil {
		return fmt.Errorf("failed to insert seat: %v", err)
	}

	// seat.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *seatRepository) UpdateSeat(ctx context.Context, seat *models.Seat) error {
	filter := bson.M{"_id": seat.ID}
	update := bson.M{"$set": seat}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update seat: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("failed to update seat: %v", err)
	}

	return nil
}

func (r *seatRepository) DeleteSeat(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid seat ID: %v", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete seat: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("failed to delete seat: %v", err)
	}

	return nil
}
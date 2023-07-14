package repositories

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShowtimeRepository interface {
	GetAllShowtimes(ctx context.Context) ([]models.Showtime, error)
	GetShowtimeByID(ctx context.Context, id string) (*models.Showtime, error)
	CreateShowtime(ctx context.Context, showtime *models.Showtime) error
	UpdateShowtime(ctx context.Context, showtime *models.Showtime) error
	DeleteShowtime(ctx context.Context, id string) error
}

type showtimeRepository struct {
	collection *mongo.Collection
}

func NewShowtimeRepository(collection *mongo.Collection) ShowtimeRepository {
	return &showtimeRepository{
		collection: collection,
	}
}

func (r *showtimeRepository) GetAllShowtimes(ctx context.Context) ([]models.Showtime, error) {
	var showtimes []models.Showtime

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve showtimes: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var showtime models.Showtime
		if err := cursor.Decode(&showtime); err != nil {
			return nil, fmt.Errorf("failed to decode showtime: %v", err)
		}
		showtimes = append(showtimes, showtime)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return showtimes, nil
}

func (r *showtimeRepository) GetShowtimeByID(ctx context.Context, id string) (*models.Showtime, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid showtime ID: %v", err)
	}

	var showtime models.Showtime
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&showtime); err != nil {
		return nil, fmt.Errorf("failed to retrieve showtime: %v", err)
	}

	return &showtime, nil
}

func (r *showtimeRepository) CreateShowtime(ctx context.Context, showtime *models.Showtime) error {
	_, err := r.collection.InsertOne(ctx, showtime)
	if err != nil {
		return fmt.Errorf("failed to insert showtime: %v", err)
	}

	// Create 64 Seats as well
	var seats []models.Seat
	for i := 1; i <= 64; i++ {
		seats = append(seats, models.Seat{
			ID:          primitive.NewObjectID(),
			ShowtimeID:  showtime.ID,
			SeatNumber:  i,
			IsAvailable: true,
		})
	}

	seatCollection := mongodb.GetCollection(mongodb.DB, "seat")

	// Convert seats slice to []interface{}
	var seatInterfaces []interface{}
	for _, seat := range seats {
		seatInterfaces = append(seatInterfaces, seat)
	}

	_, err = seatCollection.InsertMany(ctx, seatInterfaces)
	if err != nil {
		return fmt.Errorf("failed to insert seats: %v", err)
	}

	return nil
}

func (r *showtimeRepository) UpdateShowtime(ctx context.Context, showtime *models.Showtime) error {
	filter := bson.M{"_id": showtime.ID}
	update := bson.M{"$set": showtime}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update showtime: %v", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no showtime found with ID %s", showtime.ID)
	}

	return nil
}

func (r *showtimeRepository) DeleteShowtime(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid showtime ID: %v", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete showtime: %v", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no showtime found with ID %s", id)
	}

	return nil
}

package repositories

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketRepository interface {
	GetAllTickets(ctx context.Context) ([]models.Ticket, error)
	GetTicketByID(ctx context.Context, id string) (*models.Ticket, error)
	GetAllTicketsByUserID(ctx context.Context, id string) ([]models.Ticket, error)
	GetAllTicketsByUsername(ctx context.Context, username string) ([]models.Ticket, error)
	CreateTicket(ctx context.Context, ticket *models.Ticket) (models.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *models.Ticket) error
	DeleteTicket(ctx context.Context, id string) error
}

type ticketRepository struct {
	collection *mongo.Collection
}

func NewTicketRepository(collection *mongo.Collection) TicketRepository {
	return &ticketRepository{
		collection: collection,
	}
}

func (r *ticketRepository) GetAllTickets(ctx context.Context) ([]models.Ticket, error) {
	var tickets []models.Ticket

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tickets: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var ticket models.Ticket
		if err := cursor.Decode(&ticket); err != nil {
			return nil, fmt.Errorf("failed to decode ticket: %v", err)
		}
		tickets = append(tickets, ticket)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(ctx context.Context, id string) (*models.Ticket, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ticket ID: %v", err)
	}

	var ticket models.Ticket
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&ticket); err != nil {
		return nil, fmt.Errorf("failed to retrieve ticket: %v", err)
	}

	return &ticket, nil
}

func (r *ticketRepository) GetAllTicketsByUserID(ctx context.Context, id string) ([]models.Ticket, error) {
	var tickets []models.Ticket

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ticket ID: %v", err)
	}

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": objectID})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tickets: %v", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var ticket models.Ticket
		if err := cursor.Decode(&ticket); err != nil {
			return nil, fmt.Errorf("failed to decode ticket: %v", err)
		}
		tickets = append(tickets, ticket)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return tickets, nil

}

func (r *ticketRepository) GetAllTicketsByUsername(ctx context.Context, username string) ([]models.Ticket, error) {
	var tickets []models.Ticket

	userCollection := r.collection.Database().Collection("users")
	var user models.User
	if err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": user.ID})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tickets: %v", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var ticket models.Ticket
		if err := cursor.Decode(&ticket); err != nil {
			return nil, fmt.Errorf("failed to decode ticket: %v", err)
		}
		tickets = append(tickets, ticket)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return tickets, nil
}

func (r *ticketRepository) CreateTicket(ctx context.Context, ticket *models.Ticket) (models.Ticket, error) {
	ticket.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, ticket)
	if err != nil {
		return models.Ticket{}, fmt.Errorf("failed to create ticket: %v", err)
	}

	return *ticket, nil
}

func (r *ticketRepository) UpdateTicket(ctx context.Context, ticket *models.Ticket) error {
	filter := bson.M{"_id": ticket.ID}
	update := bson.M{"$set": ticket}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update ticket: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("failed to update ticket: %v", err)
	}

	return nil
}

func (r *ticketRepository) DeleteTicket(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ticket ID: %v", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete ticket: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("failed to delete ticket: %v", err)
	}

	return nil
}

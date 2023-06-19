package repositories

import (
	"context"
	"fmt"

	"github.com/Mhmdaris15/booking-movie-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode user: %v", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	var user models.User
	if err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	if _, err := r.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user}); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}

	if _, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID}); err != nil {
		return err
	}

	return nil
}
package models

import (
	"context"
	"log"
	"math/rand"

	"github.com/bxcodec/faker/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string 		   	`json:"username" bson:"username" validate:"required"`
	Email    string		   		`json:"email" bson:"email" validate:"required,email"`
	Password string	   			`json:"password" bson:"password" validate:"required"`
	Name     string	   			`json:"name" bson:"name" validate:"required"`
	Age      int	   			`json:"age" bson:"age" validate:"required"`
	Balance  float64	   		`json:"balance" bson:"balance" validate:"required"`
	// Add other fields as per your requirements
}

func SeedUser(collection *mongo.Collection) ([]User) {
	var users []User
	userCount := 5

	for i := 0; i < userCount; i++ {
		user := User{
			ID: 	 primitive.NewObjectID(),
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: "2wsx1qaz",
			Name:     faker.Name(),
			Age:      rand.Intn(100),
			Balance:  100000,
		}
		users = append(users, user)

		_, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			log.Fatal(err)
		}
	}

	return users
}


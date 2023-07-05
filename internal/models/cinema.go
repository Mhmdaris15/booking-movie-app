package models

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cinema struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	Address  string `bson:"address,omitempty" json:"address,omitempty"`
	City     string `bson:"city,omitempty" json:"city,omitempty"`
	Province string `bson:"province,omitempty" json:"province,omitempty"`
	Phone    string	`bson:"phone,omitempty" json:"phone,omitempty"`
	Email    string	`bson:"email,omitempty" json:"email,omitempty"`
}

func SeedCinema(collection *mongo.Collection) ([]Cinema) {
	var cinemas []Cinema
	cinemaCount := 5

	for i := 0; i < cinemaCount; i++ {
		cinema := Cinema{
			ID: 	 primitive.NewObjectID(),
			Name:     faker.Name(),
			Address: faker.Sentence(),
			City: faker.ChineseLastName(),
			Phone: faker.Phonenumber(),
			Province: faker.ChineseLastName(),
			Email:    faker.Email(),
		}
		cinemas = append(cinemas, cinema)

		_, err := collection.InsertOne(context.Background(), cinema)
		if err != nil {
			log.Fatal(err)
		}
	}

	return cinemas
}


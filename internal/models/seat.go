package models

import (
	"context"
	"log"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Seat struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ShowtimeID  primitive.ObjectID `bson:"showtime_id,omitempty" json:"showtime_id,omitempty"`
	SeatNumber  int
	IsAvailable bool
}

func SeedSeat(seatCollection *mongo.Collection, showtimeCollection *mongo.Collection) ([]Seat){
	var seats []Seat
	seatCount := 50

	// Retrieve all ID from showtime collection
	var showtimes []Showtime
	cursor, err  := showtimeCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.Background(), &showtimes); err != nil {
		log.Fatal(err)
	}

	showtimeIDs := []primitive.ObjectID{}

	for _, showtime := range showtimes {
		showtimeIDs = append(showtimeIDs, showtime.ID)
	}

	for i := 0; i < seatCount; i++ {
		seat := Seat{
			ID:          primitive.NewObjectID(),
			ShowtimeID:  showtimeIDs[rand.Intn(len(showtimeIDs))],
			SeatNumber:  rand.Intn(64 - 1 + 1) + 1,
			IsAvailable: true,
		}
		seats = append(seats, seat)

		_, err := seatCollection.InsertOne(context.Background(), seat)
		if err != nil {
			log.Fatal(err)
		}
	}

	return seats
}
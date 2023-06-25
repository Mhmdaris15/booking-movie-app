package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	MovieID         primitive.ObjectID  `bson:"movie_id,omitempty" json:"movie_id,omitempty" validate:"required"`
	UserID          primitive.ObjectID  `bson:"user_id,omitempty" json:"user_id,omitempty" validate:"required"`
	ShowtimeID      primitive.ObjectID  `bson:"showtime_id,omitempty" json:"showtime_id,omitempty" validate:"required"`
	SeatID          []primitive.ObjectID  `bson:"seat_id,omitempty" json:"seat_id,omitempty" validate:"required"`
	TransactionDate string `bson:"transactionDate,omitempty" json:"transactionDate date" validate:"required"`
	TotalPrice      float64 `bson:"totalPrice,omitempty" json:"totalPrice,omitempty" validate:"required"`
	// Add other fields as per your requirements
}

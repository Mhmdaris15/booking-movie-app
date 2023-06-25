package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string 		  `json:"title" bson:"title" validate:"required"`
	Description string 		  `json:"description" bson:"description" validate:"required"`
	AgeRating   float64 	  `json:"age_rating" bson:"age_rating" validate:"required"`
	PosterURL   string 		  `json:"poster_url" bson:"poster_url" validate:"required"`
	TicketPrice float64		  `json:"ticket_price" bson:"ticket_price" validate:"required"`
	// Add other fields as per your requirements
}

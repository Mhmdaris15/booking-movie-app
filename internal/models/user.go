package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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

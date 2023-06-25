package response

import (
	"github.com/Mhmdaris15/booking-movie-app/internal/models"
)

type UserResponse struct {
	Status int    `json:"status"`
	Message string `json:"message"`
	Data    []models.User `json:"data"`
}
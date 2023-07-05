package response

import (
	"github.com/Mhmdaris15/booking-movie-app/internal/models"
)

type Response struct {
	Data  models.Movie `json:"movie_data"`
	Image string `json:"movie_image"`
}

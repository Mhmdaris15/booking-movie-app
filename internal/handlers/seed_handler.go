package handlers

import (
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongodb"
	"github.com/gin-gonic/gin"
)

func SeedingDatabase(c *gin.Context) {
	db := mongodb.ConnectDB()
	user, cinemas, showtimes, seats := mongodb.SeedingDatabase(db)
	mongodb.DisconnectDB(db)
	c.JSON(200, gin.H{
		"message": "Seeding Database Success",
		"user": user,
		"cinema": cinemas,
		"showtime": showtimes,
		"seat":seats,
	})
	
}
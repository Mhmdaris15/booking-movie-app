package handlers

import (
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongo"
	"github.com/gin-gonic/gin"
)

func SeedingDatabase(c *gin.Context) {
	db := mongo.ConnectDB()
	user, cinemas, showtimes, seats := mongo.SeedingDatabase(db)
	mongo.DisconnectDB(db)
	c.JSON(200, gin.H{
		"message": "Seeding Database Success",
		"user": user,
		"cinema": cinemas,
		"showtime": showtimes,
		"seat":seats,
	})
	
}
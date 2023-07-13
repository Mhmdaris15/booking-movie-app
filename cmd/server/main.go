package main

import (
	"log"

	"github.com/Mhmdaris15/booking-movie-app/internal/configs"
	"github.com/Mhmdaris15/booking-movie-app/internal/routes"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to MongoDB
	mongodb.ConnectDB()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3005"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Expose-Headers"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	routes.SetupRoutes(router)

	if err := router.Run(configs.Port()); err != nil {
		log.Fatal(err.Error())
	}
}

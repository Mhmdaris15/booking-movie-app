package main

import (
	"log"

	"github.com/Mhmdaris15/booking-movie-app/internal/routes"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
        // Load environment variables from .env file
        // err := godotenv.Load()
        // if err != nil {
        //         log.Fatal("Error loading .env file")
        // }

        // Connect to MongoDB
        if err := mongo.Connect("mongodb://localhost:27017", "moviedb"); err != nil {
                log.Fatal("Error connecting to MongoDB", err.Error())
        }


        router := gin.Default()

        config := cors.DefaultConfig()
  	config.AllowAllOrigins = true

        router.Use(cors.New(config))

        routes.SetupRoutes(router)

        if err := router.Run(":3000"); err != nil {
                log.Fatal(err.Error())
        }
}
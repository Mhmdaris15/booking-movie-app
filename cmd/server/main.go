package main

import (
	"log"

	"github.com/Mhmdaris15/booking-movie-app/internal/configs"
	"github.com/Mhmdaris15/booking-movie-app/internal/routes"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
        
        // Connect to MongoDB
	mongo.ConnectDB()

        router := gin.Default()

        config := cors.DefaultConfig()
  	config.AllowAllOrigins = true

        router.Use(cors.New(config))

        routes.SetupRoutes(router)

        if err := router.Run(configs.Port()); err != nil {
                log.Fatal(err.Error())
        }
}
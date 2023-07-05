package routes

import (
	"github.com/Mhmdaris15/booking-movie-app/internal/handlers"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongo"
	"github.com/Mhmdaris15/booking-movie-app/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){

	// Get Collections
	userCollection := mongo.GetCollection(mongo.DB, "users")
	movieCollection := mongo.GetCollection(mongo.DB, "movie")
	ticketCollection := mongo.GetCollection(mongo.DB, "ticket")
	cinemaCollection := mongo.GetCollection(mongo.DB, "cinema")
	seatCollection := mongo.GetCollection(mongo.DB, "seat")
	showtimeCollection := mongo.GetCollection(mongo.DB, "showtime")

	
	// Setup Repositories
	userRepository := repositories.NewUserRepository(userCollection)
	movieRepository := repositories.NewMovieRepository(movieCollection)
	ticketRepository := repositories.NewTicketRepository(ticketCollection)
	cinemaRepository := repositories.NewCinemaRepository(cinemaCollection)
	seatRepository := repositories.NewSeatRepository(seatCollection)
	showtimeRepository := repositories.NewShowtimeRepository(showtimeCollection)

	// Setup handlers
	userHandler := handlers.NewUserHandler(services.NewUserService(userRepository))
	movieHandler := handlers.NewMovieHandler(services.NewMovieService(movieRepository))
	ticketHandler := handlers.NewTicketHandler(services.NewTicketService(ticketRepository))
	cinemaHandler := handlers.NewCinemaHandler(services.NewCinemaService(cinemaRepository))
	seatHandler := handlers.NewSeatHandler(services.NewSeatService(seatRepository))
	showtimeHandler := handlers.NewShowtimeHandler(services.NewShowtimeService(showtimeRepository))

	// Setup routes

	// User routes
	router.GET("/users", userHandler.GetAllUsers)
	// router.GET("/users/:id", userHandler.GetUserByID)
	// router.GET("/users/:id/tickets", userHandler.GetUserTickets)
	// router.GET("/users/:email", userHandler.GetUserByEmail)
	router.GET("/users/:username", userHandler.GetUserByUsername)
	router.POST("/users", userHandler.CreateUser)
	// router.PUT("/users/:id", userHandler.UpdateUser)
	router.PUT("/users/:username", userHandler.UpdateUserByUsername)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	// Movie routes
	router.GET("/movies", movieHandler.GetAllMovies)
	router.GET("/movies/:id", movieHandler.GetMovieByID)
	router.POST("/movies", movieHandler.CreateMovie)
	router.PUT("/movies/:id", movieHandler.UpdateMovie)
	router.DELETE("/movies/:id", movieHandler.DeleteMovie)

	// Ticket routes
	router.GET("/tickets", ticketHandler.GetAllTickets)
	router.GET("/tickets/:id", ticketHandler.GetTicketByID)
	router.POST("/tickets", ticketHandler.CreateTicket)
	router.PUT("/tickets/:id", ticketHandler.UpdateTicket)
	router.DELETE("/tickets/:id", ticketHandler.DeleteTicket)
	
	// Cinema routes
	router.GET("/cinemas", cinemaHandler.GetAllCinemas)
	router.GET("/cinemas/:id", cinemaHandler.GetCinemaByID)
	router.POST("/cinemas", cinemaHandler.CreateCinema)
	router.PUT("/cinemas/:id", cinemaHandler.UpdateCinema)
	router.DELETE("/cinemas/:id", cinemaHandler.DeleteCinema)

	// Seat routes
	router.GET("/seats", seatHandler.GetAllSeats)
	router.GET("/seats/:id", seatHandler.GetSeatByID)
	router.POST("/seats", seatHandler.CreateSeat)
	router.PUT("/seats/:id", seatHandler.UpdateSeat)
	router.DELETE("/seats/:id", seatHandler.DeleteSeat)

	// Showtime routes
	router.GET("/showtimes", showtimeHandler.GetAllShowtimes)
	router.GET("/showtimes/:id", showtimeHandler.GetShowtimeByID)
	router.POST("/showtimes", showtimeHandler.CreateShowtime)
	router.PUT("/showtimes/:id", showtimeHandler.UpdateShowtime)
	router.DELETE("/showtimes/:id", showtimeHandler.DeleteShowtime)

	// Authentication
	router.POST("/register", utils.Signup)
	router.POST("/login", utils.Login)
	
	// Seeding Database
	router.GET("/seed", handlers.SeedingDatabase)

	// defer db.Close()
}


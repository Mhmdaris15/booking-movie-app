package routes

import (
	"github.com/Mhmdaris15/booking-movie-app/internal/handlers"
	"github.com/Mhmdaris15/booking-movie-app/internal/repositories"
	"github.com/Mhmdaris15/booking-movie-app/internal/services"
	"github.com/Mhmdaris15/booking-movie-app/pkg/database/mongodb"
	"github.com/Mhmdaris15/booking-movie-app/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine){

	// Get Collections
	userCollection := mongodb.GetCollection(mongodb.DB, "users")
	movieCollection := mongodb.GetCollection(mongodb.DB, "movie")
	ticketCollection := mongodb.GetCollection(mongodb.DB, "ticket")
	cinemaCollection := mongodb.GetCollection(mongodb.DB, "cinema")
	seatCollection := mongodb.GetCollection(mongodb.DB, "seat")
	showtimeCollection := mongodb.GetCollection(mongodb.DB, "showtime")

	
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
	userGroup := router.Group("/users")
	userGroup.Use(utils.AuthMiddleware())
	{
		userGroup.PUT("/users/:username", userHandler.UpdateUserByUsername)
		userGroup.DELETE("/users/:id", userHandler.DeleteUser)
	}

	// Movie routes
	router.GET("/movies", movieHandler.GetAllMovies)
	router.GET("/movies/:id", movieHandler.GetMovieByID)
	
	movieGroup := router.Group("/movies")
	movieGroup.Use(utils.AuthMiddleware())
	{
		movieGroup.POST("/movies", movieHandler.CreateMovie)
		movieGroup.PUT("/movies/:id", movieHandler.UpdateMovie)
		movieGroup.DELETE("/movies/:id", movieHandler.DeleteMovie)
	}

	// Ticket routes

	ticketGroup := router.Group("/tickets")
	ticketGroup.Use(utils.AuthMiddleware())
	{
		ticketGroup.GET("/user/:id", ticketHandler.GetAllTicketsByUserID)
		ticketGroup.GET("/:id", ticketHandler.GetTicketByID)
		ticketGroup.GET("/", ticketHandler.GetAllTickets)
		ticketGroup.POST("/", ticketHandler.CreateTicket)
		ticketGroup.PUT("/:id", ticketHandler.UpdateTicket)
		ticketGroup.DELETE("/:id", ticketHandler.DeleteTicket)
	}

	
	// Cinema routes
	router.GET("/cinemas", cinemaHandler.GetAllCinemas)
	router.GET("/cinemas/:id", cinemaHandler.GetCinemaByID)

	cinemaGroup := router.Group("/cinemas")
	cinemaGroup.Use(utils.AuthMiddleware())
	{
		cinemaGroup.POST("/", cinemaHandler.CreateCinema)
		cinemaGroup.PUT("/:id", cinemaHandler.UpdateCinema)
		cinemaGroup.DELETE("/:id", cinemaHandler.DeleteCinema)
	}

	// Seat routes
	router.GET("/seats", seatHandler.GetAllSeats)

	seatGroup := router.Group("/seats")
	seatGroup.Use(utils.AuthMiddleware())
	{
		seatGroup.GET("/seats/:id", seatHandler.GetSeatByID)
		seatGroup.POST("/seats", seatHandler.CreateSeat)
		seatGroup.PUT("/seats/:id", seatHandler.UpdateSeat)
		seatGroup.DELETE("/seats/:id", seatHandler.DeleteSeat)
	}

	// Showtime routes
	router.GET("/showtimes", showtimeHandler.GetAllShowtimes)
	router.GET("/showtimes/:id", showtimeHandler.GetShowtimeByID)
	
	showtimeGroup := router.Group("/showtimes")
	showtimeGroup.Use(utils.AuthMiddleware())
	{
		showtimeGroup.POST("/showtimes", showtimeHandler.CreateShowtime)
		showtimeGroup.PUT("/showtimes/:id", showtimeHandler.UpdateShowtime)
		showtimeGroup.DELETE("/showtimes/:id", showtimeHandler.DeleteShowtime)
	}

	// Authentication
	router.POST("/register", utils.Signup)
	router.POST("/login", utils.Login)


	router.GET("/protected", utils.AuthMiddleware(), utils.ProtectedHandler)
	// Seeding Database
	router.GET("/seed", handlers.SeedingDatabase)

	// defer db.Close()
}


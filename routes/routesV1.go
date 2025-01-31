package routes

import (
	"github.com/LuisSilva7/theatre-reservation-api/handlers"
	"github.com/LuisSilva7/theatre-reservation-api/services"
	"github.com/LuisSilva7/theatre-reservation-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize services
	authService := services.NewAuthService(db)
	showService := services.NewShowService(db)
	showtimeService := services.NewShowtimeService(db)
	reservationService := services.NewReservationService(db)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	showHandler := handlers.NewShowHandler(showService)
	showtimeHandler := handlers.NewShowtimeHandler(showtimeService)
	reservationHandler := handlers.NewReservationHandler(reservationService)

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
		public.GET("/shows", showHandler.GetShows)
		public.GET("/shows/:showID", showHandler.GetShowByID)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(utils.AuthMiddleware())
	{
		protected.GET("/user/reservations", reservationHandler.GetUserReservations)
		protected.POST("/user/reservations", reservationHandler.CreateReservation)
		protected.DELETE("/user/reservations/:reservationID", reservationHandler.CancelReservation)
		protected.GET("/showtimes/:showtimeID/seats", reservationHandler.GetAvailableSeats)
	}

	// Admin routes
	admin := router.Group("/api/v1")
	admin.Use(utils.AuthMiddleware(), utils.AdminMiddleware())
	{
		admin.POST("/shows", showHandler.AddShow)
		admin.DELETE("/shows/:showID", showHandler.DeleteShow)
		admin.GET("/reservations", reservationHandler.GetAllReservations)
		admin.POST("/users/:userID/promote", authHandler.PromoteToAdmin)
		admin.POST("/showtimes", showtimeHandler.AddShowtime)
		admin.DELETE("/showtimes/:showtimeID", showtimeHandler.DeleteShowtime)
		admin.GET("/shows/:showID/report", showHandler.GetReport)
	}

	return router
}

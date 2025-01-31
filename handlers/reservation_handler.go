package handlers

import (
	"net/http"

	"github.com/LuisSilva7/theatre-reservation-api/models"
	"github.com/LuisSilva7/theatre-reservation-api/services"
	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	ReservationService *services.ReservationService
}

func NewReservationHandler(reservationService *services.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		ReservationService: reservationService,
	}
}

func (rh *ReservationHandler) GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("showtimeID")

	seats, err := rh.ReservationService.GetAvailableSeats(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seats)
}

func (rh *ReservationHandler) CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userID format"})
		return
	}

	reservation.UserID = userIDUint

	newReservation, err := rh.ReservationService.CreateReservation(reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newReservation)
}

func (rh *ReservationHandler) GetUserReservations(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	reservations, err := rh.ReservationService.GetUserReservations(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func (rh *ReservationHandler) CancelReservation(c *gin.Context) {
	reservationID := c.Param("reservationId")
	if err := rh.ReservationService.CancelReservation(reservationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation cancelled successfully"})
}

func (rh *ReservationHandler) GetAllReservations(c *gin.Context) {
	reservations, err := rh.ReservationService.GetAllReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

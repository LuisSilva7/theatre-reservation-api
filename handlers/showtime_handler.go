
package handlers

import (
	"net/http"

	"github.com/LuisSilva7/theatre-reservation-api/models"
	"github.com/LuisSilva7/theatre-reservation-api/services"
	"github.com/gin-gonic/gin"
)

type ShowtimeHandler struct {
	ShowtimeService *services.ShowtimeService
}

func NewShowtimeHandler(showtimeService *services.ShowtimeService) *ShowtimeHandler {
	return &ShowtimeHandler{
		ShowtimeService: showtimeService,
	}
}

func (sh *ShowtimeHandler) AddShowtime(c *gin.Context) {
	var showtime models.Showtime
	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedShowtime, err := sh.ShowtimeService.AddShowtime(showtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, addedShowtime)
}

func (sh *ShowtimeHandler) GetShowtimes(c *gin.Context) {
	showID := c.Param("showID")
	showtimes, err := sh.ShowtimeService.GetShowtimes(showID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, showtimes)
}

func (sh *ShowtimeHandler) DeleteShowtime(c *gin.Context) {
	showtimeID := c.Param("showtimeID")
	err := sh.ShowtimeService.DeleteShowtime(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "showtime deleted successfully"})
}

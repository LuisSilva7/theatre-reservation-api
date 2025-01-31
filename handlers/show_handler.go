package handlers

import (
	"net/http"
	"time"

	"github.com/LuisSilva7/theatre-reservation-api/models"
	"github.com/LuisSilva7/theatre-reservation-api/services"
	"github.com/gin-gonic/gin"
)

type ShowHandler struct {
	ShowService *services.ShowService
}

func NewShowHandler(showService *services.ShowService) *ShowHandler {
	return &ShowHandler{
		ShowService: showService,
	}
}

type AddShowRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Duration    int    `json:"duration"`
	ReleaseDate string `json:"realease_date"`
}

func (sh *ShowHandler) AddShow(c *gin.Context) {
	var req AddShowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid realease date"})
		return
	}

	show := models.Show{
		Name:        req.Name,
		Description: req.Description,
		Genre:       models.Genre(req.Genre),
		Duration:    req.Duration,
		ReleaseDate: releaseDate,
	}

	addedShow, err := sh.ShowService.AddShow(show)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, addedShow)
}

func (sh *ShowHandler) GetShows(c *gin.Context) {
	shows, err := sh.ShowService.GetShows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shows)
}

func (sh *ShowHandler) GetShowByID(c *gin.Context) {
	showID := c.Param("showID")
	show, err := sh.ShowService.GetShowByID(showID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, show)
}

func (sh *ShowHandler) DeleteShow(c *gin.Context) {
	showID := c.Param("showID")
	err := sh.ShowService.DeleteShow(showID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "show deleted successfully"})
}

func (sh *ShowHandler) GetReport(c *gin.Context) {
	// TODO - report
}

package handlers

import (
	"net/http"

	"github.com/LuisSilva7/theatre-reservation-api/models"
	"github.com/LuisSilva7/theatre-reservation-api/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (ac *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := ac.AuthService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (ac *AuthHandler) Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := ac.AuthService.Login(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (ac *AuthHandler) PromoteToAdmin(c *gin.Context) {
	userID := c.Param("userID")
	if err := ac.AuthService.PromoteToAdmin(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}

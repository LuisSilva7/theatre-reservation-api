package main

import (
	"log"

	"github.com/LuisSilva7/theatre-reservation-api/config"
	"github.com/LuisSilva7/theatre-reservation-api/models"
	"github.com/LuisSilva7/theatre-reservation-api/routes"
	"github.com/LuisSilva7/theatre-reservation-api/utils"
	"gorm.io/gorm"
)

func createAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Where("email = ?", "admin@email.com").Count(&count)

	if count == 0 {
		hashPassword, err := utils.HashPassword("admin123")
		if err != nil {
			return err
		}

		user := models.User{
			FirstName: "Luís",
			LastName: "Silva",
			Email:    "admin@example.com",
			Password: hashPassword,
			Role:     models.AdminRole,
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := createAdmin(db); err != nil {
		log.Fatalf("failed to create admin")
	}

	router := routes.SetupRouter(db)
	router.Run(":8888")
}

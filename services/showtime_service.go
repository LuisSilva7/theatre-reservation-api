package services

import (
	"github.com/LuisSilva7/theatre-reservation-api/models"
	"gorm.io/gorm"
)

type ShowtimeService struct {
	DB *gorm.DB
}

func NewShowtimeService(db *gorm.DB) *ShowtimeService {
	return &ShowtimeService{
		DB: db,
	}
}

func (ss *ShowtimeService) AddShowtime(showtime models.Showtime) (*models.Showtime, error) {
	if err := ss.DB.Create(&showtime).Error; err != nil {
		return nil, err
	}
	return &showtime, nil
}

func (ss *ShowtimeService) GetShowtimes(showID string) ([]*models.Showtime, error) {
	var showtimes []*models.Showtime
	if err := ss.DB.Where("show_id = ?", showID).Find(&showtimes).Error; err != nil {
		return nil, err
	}
	return showtimes, nil
}

func (ss *ShowtimeService) DeleteShowtime(showtimeID string) error {
	if err := ss.DB.Where("id = ?", showtimeID).Delete(&models.Showtime{}).Error; err != nil {
		return err
	}
	return nil
}

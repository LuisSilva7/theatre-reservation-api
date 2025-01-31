package services

import (
	"github.com/LuisSilva7/theatre-reservation-api/models"
	"gorm.io/gorm"
)

type ShowService struct {
	DB *gorm.DB
}

func NewShowService(db *gorm.DB) *ShowService {
	return &ShowService{
		DB: db,
	}
}

type Report struct {
}

func (ss *ShowService) AddShow(show models.Show) (*models.Show, error) {
	if err := ss.DB.Create(&show).Error; err != nil {
		return nil, err
	}

	return &show, nil
}

func (ss *ShowService) GetShows() ([]*models.Show, error) {
	var shows []*models.Show
	if err := ss.DB.Find(&shows).Error; err != nil {
		return nil, err
	}

	return shows, nil
}

func (ss *ShowService) GetShowByID(showID string) (*models.Show, error) {
	var show models.Show
	if err := ss.DB.Where("id = ?", showID).First(&show).Error; err != nil {
		return nil, err
	}

	return &show, nil
}

func (ss *ShowService) DeleteShow(showID string) error {
	if err := ss.DB.Where("id = ?", showID).Delete(&models.Show{}).Error; err != nil {
		return nil
	}

	return nil
}

func (ss *ShowService) GetReport() (Report, error) {
	// TODO - report
	return Report{}, nil
}

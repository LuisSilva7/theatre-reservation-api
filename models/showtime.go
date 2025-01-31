package models

import "time"

type Showtime struct {
	ID             uint      `gorm:"primary_key;" json:"id"`
	MovieID        uint      `gorm:"not null;" json:"show_id"`
	StartTime      time.Time `gorm:"not null;" json:"start_time"`
	EndTime        time.Time `gorm:"not null;" json:"end_time"`
	AvailableSeats int       `gorm:"not null;" json:"available_seats"`
	Price          float64   `gorm:"not null;" json:"price"`
	CreatedAt      time.Time `json:"createdAt" json:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" json:"updated_at"`
}

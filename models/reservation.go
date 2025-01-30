package models

import "time"

type Reservation struct {
	ID              uint      `gorm:"primary_key;" json:"id"`
	UserID          uint      `gorm:"not null;" json:"user_id"`
	ShowtimeID      uint      `gorm:"not null;" json:"showtime_id"`
	SeatNumbers     string    `gorm:"not null;" json:"seat_numbers"`
	ReservationTime time.Time `gorm:"not null;" json:"reservation_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

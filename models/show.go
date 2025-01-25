package main

import "time"

type Genre string

const (
	DramaGenre          Genre = "Drama"
	ComedyGenre         Genre = "Comedy"
	MusicalGenre        Genre = "Musical"
	ScienceFictionGenre Genre = "Science Fiction"
)

type Show struct {
	ID          uint      `gorm:"primary_key;" json:"id"`
	Name        string    `gorm:"not null;" json:"name"`
	Description string    `gorm:"not null;" json:"description"`
	Genre       Genre     `gorm:"not null;" json:"genre"`
	Duration    int       `json:"duration"`
	ReleaseDate time.Time `json:"realease_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

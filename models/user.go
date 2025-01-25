package models

import "time"

type Role int

const (
	AdminRole Role = iota
	RegularUserRole
)

type User struct {
	ID        uint      `gorm:"primary_key;" json:"id"`
	FirstName string    `gorm:"not null;" json:"first_name"`
	LastName  string    `gorm:"not null;" json="last_name"`
	Email     string    `gorm:"unique;not null;" json="email"`
	Password  string    `gorm:"not null;" json="password"`
	Role      Role      `gorm:"not null;" json="role"`
	CreatedAt time.Time `json="created_at"`
	UpdatedAt time.Time `json="updated_at"`
}

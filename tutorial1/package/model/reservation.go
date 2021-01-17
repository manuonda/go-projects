package model

import (
	"time"
	"gorm.io/gorm"
)

type Reservation struct{
	gorm.Model

	// Book Book
	BookID int
	TelephoneNumber string
	DNI string
	ReturnedAt *time.Time
}
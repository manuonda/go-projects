package models

import "time"

type POST struct {
	ID        uint      `json:"id"      gorm:"primaryKey"`
	Title     string    `json:"title"   gorm:"title"`
	Content   string    `json:"content" gorm:"content"`
	CreatedAt time.Time `json:"created_at"`
}

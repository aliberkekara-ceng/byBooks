package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" binding:"required" gorm:"not null"`
	Author    string    `json:"author" binding:"required" gorm:"not null"`
	Year      int       `json:"year" binding:"required,gt=0" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

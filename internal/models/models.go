package models

import "time"

type Application struct {
	ID        uint `gorm:"primaryKey"`
	Company   string
	Position  string
	Link      string
	Status    string
	AppliedAt time.Time
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint // связь с пользователем
}

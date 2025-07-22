package model

import "github.com/google/uuid"

type Subscription struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Service   string    `json:"service_name" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	StartDate string    `json:"start_date" binding:"required"`
	EndDate   *string   `json:"end_date"`
}

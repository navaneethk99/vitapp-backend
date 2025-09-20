package models

import "time"

type Carpool struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	PickupLocation string    `json:"pickupLocation"`
	DropLocation   string    `json:"dropLocation"`
	Datetime       time.Time `json:"datetime"`
}

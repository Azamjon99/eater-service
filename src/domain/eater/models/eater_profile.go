package models

import "time"

type EaterProfile struct {
	EaterID     string    `json:"eater_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

package models

import "time"

type Eater struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	PhoneNumber  string    `json:"phone_number"`
	PasswordHash string    `json:"password_hash"`
	PasswordSalt string    `json:"password_salt"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

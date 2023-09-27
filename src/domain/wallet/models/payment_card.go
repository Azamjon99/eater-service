package models

import (
	"time"
)

type PaymentCard struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	EaterID    string    `json:"eater_id"`
	Number     string    `json:"number"`
	CardToken  string    `json:"card_token"`
	IsVerifed  bool      `json:"is_verifed"`
	CreatedAt  time.Time `json:"created_at"`
}

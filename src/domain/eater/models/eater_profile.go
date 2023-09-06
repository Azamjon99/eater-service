package models

import "time"

type EaterProfile struct {
	EaterID string		    `json:"eater_id" gorm:"index:idx_sms_code_by_eater"`		
	Name string				`json:"password_hash"`
	PhoneNumber string		`json:"phone_number"`
	ImageUrl string			 `json:"image_url"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
	}
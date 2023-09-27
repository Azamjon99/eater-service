package models

import "time"

type Order struct {
	ID            string          `json:"id" gorm:"primaryKey"`
	EaterID       string          `json:"eater_id"`
	Instruction   string          `json:"instruction"`
	RestaurantID  string          `json:"restaurant_id"` // restaurant entity id
	RestaurantInfo    *RestaurantInfo `json:"restaurant"`
	DeliveryInfo      *DeliveryInfo   `json:"delivery"`
	PaymentInfo       *PaymentInfo    `json:"payment"`
	OrderItem         []*OrderItem    `json:"items"`
	Status        string          `json:"status"`
	PaymentStatus string          `json:"payment_status"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"` 
}

type OrderItem struct {
	ID         string    `json:"id"`
	ProductID  string    `json:"product_id"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

type RestaurantInfo struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type AddressInfo struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type DeliveryInfo struct {
	Address *AddressInfo `json:"address"`
	Time    string       `json:"time"`
	Notes   string       `json:"notes"`
}

type PaymentInfo struct {
	Method        string `json:"method"`
	CardID        string `json:"card_id"`
	DeliveryPrice int    `json:"delivery_price"`
	TotalPrice    int    `json:"total_price"`
}

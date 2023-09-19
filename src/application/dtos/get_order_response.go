package dtos

import "eater-service/src/domain/order/models"


type GetOrderResponse struct {
	Order *models.Order `json:"order"`
}

func NewGetOrderResponse(order *models.Order) *GetOrderResponse {
	return &GetOrderResponse{
		Order: order,
	}
}

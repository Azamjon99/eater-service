package dtos

import "eater-service/src/domain/order/models"


type ListOrderResponse struct {
	Order []*models.Order `json:"order"`
}

func NewListOrderResponse(order []*models.Order) *ListOrderResponse {
	return &ListOrderResponse{
		Order: order,
	}
}

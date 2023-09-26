package dtos

import "github.com/Azamjon99/eater-service/src/domain/rating/models"

type ListDeliveryRatingResponse struct {
	DeliveryRating []*models.DeliveryRating `json:"delivery_rating"`
}

func NewListDeliveryRatingResponse(deliveryRating []*models.DeliveryRating) *ListDeliveryRatingResponse {
	return &ListDeliveryRatingResponse{
		DeliveryRating: deliveryRating,
	}
}

package dtos

import "eater-service/src/domain/rating/models"


type GetDeliveryRatingResponse struct {
	DeliveryRating *models.DeliveryRating `json:"delivery_rating"`
}

func NewGetDeliveryRatingResponse(deliveryRating *models.DeliveryRating) *GetDeliveryRatingResponse {
	return &GetDeliveryRatingResponse{
		DeliveryRating: deliveryRating,
	}
}

package dtos

import "eater-service/src/domain/rating/models"



type ListRestaurantRatingResponse struct {
	RestaurantRating []*models.RestaurantRating `json:"restaurant_rating"`
}

func NewListRestaurantRatingResponse(restaurantRating []*models.RestaurantRating) *ListRestaurantRatingResponse {
	return &ListRestaurantRatingResponse{
		RestaurantRating: restaurantRating,
	}
}

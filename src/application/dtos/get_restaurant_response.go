package dtos

import "eater-service/src/domain/rating/models"


type GetRestaurantRatingResponse struct {
	RestaurantRating *models.RestaurantRating `json:"restaurant_rating"`
}

func NewGetRestaurantRatingResponse(restaurantRating *models.RestaurantRating) *GetRestaurantRatingResponse {
	return &GetRestaurantRatingResponse{
		RestaurantRating: restaurantRating,
	}
}

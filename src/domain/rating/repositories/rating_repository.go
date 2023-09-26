package repositories

import (
	"context"
	"github.com/Azamjon99/eater-service/src/domain/rating/models"
)

type RatingRepository interface {
	RateRestaurant(ctx context.Context, restaurantRatings *models.RestaurantRating) error
	UpdateRestaurantRating(ctx context.Context, restaurantRatings *models.RestaurantRating) error
	ListRestaurantRatingEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.RestaurantRating, error)
	RateDelivery(ctx context.Context, deliveryRatings *models.DeliveryRating) error
	UpdateDeliveryRating(ctx context.Context, deliveryRatings *models.DeliveryRating) error
	GetDeliveryRatingByOrderId(ctx context.Context, orderId string) (*models.DeliveryRating, error)
	ListDeliveryRatingByEaterId(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
}

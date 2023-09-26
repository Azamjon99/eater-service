package rating

import (
	"context"

	"eater-service/src/domain/rating/models"
	"eater-service/src/domain/rating/repositories"
	"eater-service/src/infrastructure/repositories/utils"
	"gorm.io/gorm"
)

const (
	tableDeliveryRating   = "rating.delivery_rating"
	tableRestaurantRating = "rating.restaurant_rating"
)

type ratingRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.RatingRepository {
	return &ratingRepoImpl{
		db: db,
	}
}



func (r *ratingRepoImpl) RateRestaurant(ctx context.Context, restaurantRatings *models.RestaurantRating) error {
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Create(restaurantRatings)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) UpdateRestaurantRating(ctx context.Context, restaurantRatings *models.RestaurantRating) error {
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Save(restaurantRatings)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) ListRestaurantRatingEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.RestaurantRating, error) {
	var restaurantRatings []*models.RestaurantRating
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Where("eater_id = ?", eaterID)
	result.Scopes(utils.SortByCreatedAt(page, pageSize), utils.Sort(sort)).Find(&restaurantRatings)

	if result.Error != nil {
		return nil, result.Error
	}
	return restaurantRatings, nil
}

func (r *ratingRepoImpl) RateDelivery(ctx context.Context, deliveryRatings *models.DeliveryRating) error {
	result := r.db.WithContext(ctx).Table(tableDeliveryRating).Create(deliveryRatings)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) UpdateDeliveryRating(ctx context.Context, deliveryRatings *models.DeliveryRating) error {
	result := r.db.WithContext(ctx).Table(tableDeliveryRating).Save(deliveryRatings)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ratingRepoImpl) GetDeliveryRatingByOrderId(ctx context.Context, orderId string) (*models.DeliveryRating, error) {
	var deliveryRatings *models.DeliveryRating
	result := r.db.WithContext(ctx).Table(tableDeliveryRating).First(&deliveryRatings, "order_id = ?", orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	return deliveryRatings, nil
}

func (r *ratingRepoImpl) ListDeliveryRatingByEaterId(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error) {

	var deliveryRatings []*models.DeliveryRating
	result := r.db.WithContext(ctx).Table(tableDeliveryRating).Where("eater_id = ?", eaterID)

	if result.Error != nil {
		return nil, result.Error
	}
	return deliveryRatings, nil
}

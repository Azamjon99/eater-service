package services

import (
	"context"
	"eater-service/src/domain/rating/models"
	"eater-service/src/domain/rating/repositories"
	"eater-service/src/infrastructure/rand"
	"time"
)

type RatingService interface {
	RateRestaurant(ctx context.Context, eaterID string, restaurantID string, rating int, comment string) (*models.RestaurantRating, error)
	UpdateRestaurantRating(ctx context.Context, restaurantRatingId string, rating int, comment string) (*models.RestaurantRating, error)
	ListRestaurantRatingEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.RestaurantRating, error)
	RateDelivery(ctx context.Context, eaterID string, orderId string, rating int, comment string) (*models.DeliveryRating, error)
	UpdateDeliveryRating(ctx context.Context, deliveryRatingId string, rating int, comment string) (*models.DeliveryRating, error)
	GetDeliveryRatingByOrderId(ctx context.Context, orderId string) (*models.DeliveryRating, error)
	ListDeliveryRatingByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.DeliveryRating, error)
}
type ratingSvcImpl struct {
	ratingRepo repositories.RatingRepository
}

func NewWalletService(ratingRepo repositories.RatingRepository) RatingService {
	return &ratingSvcImpl{
		ratingRepo: ratingRepo,
	}
}

func (s *ratingSvcImpl) RateRestaurant(ctx context.Context, eaterID string, restaurantID string, rating int, comment string) (*models.RestaurantRating, error) {

	restaurantRating := &models.RestaurantRating{
		ID:           rand.UUID(),
		EaterID:      eaterID,
		RestaurantID: restaurantID,
		Rating:       rating,
		Comment:      comment,
		CreatedAt:    time.Now().UTC(),
	}

	err := s.ratingRepo.RateRestaurant(ctx, restaurantRating); 
			
	if err != nil {
		return nil, err
	}

	return restaurantRating, nil

}

func (s *ratingSvcImpl) UpdateRestaurantRating(ctx context.Context, restaurantRatingId string, rating int, comment string) (*models.RestaurantRating, error) {

	restaurantRating := &models.RestaurantRating{
		ID:        restaurantRatingId,
		Rating:    rating,
		Comment:   comment,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.UpdateRestaurantRating(ctx, restaurantRating)
	if err != nil {
		return nil, err
	}

	return restaurantRating, nil

}

func (s *ratingSvcImpl) ListRestaurantRatingEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.RestaurantRating, error) {

	restaurantRatings, err := s.ratingRepo.ListRestaurantRatingEaterId(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return restaurantRatings, nil
}

func (s *ratingSvcImpl) RateDelivery(ctx context.Context, eaterID string, orderId string, rating int, comment string) (*models.DeliveryRating, error) {

	deliveryRating := &models.DeliveryRating{
		ID:        rand.UUID(),
		EaterID:   eaterID,
		OrderID:   orderId,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.RateDelivery(ctx, deliveryRating) 
		
	if err != nil {
		return nil, err
	}

	return deliveryRating, nil

}

func (s *ratingSvcImpl) UpdateDeliveryRating(ctx context.Context, deliveryRatingId string, rating int, comment string) (*models.DeliveryRating, error) {

	deliveryRating := &models.DeliveryRating{
		ID:        deliveryRatingId,
		Rating:    rating,
		Comment:   comment,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.UpdateDeliveryRating(ctx, deliveryRating);
	if err != nil {
		return nil, err
	}

	return deliveryRating, nil

}

func (s *ratingSvcImpl) GetDeliveryRatingByOrderId(ctx context.Context, orderId string) (*models.DeliveryRating, error) {

	deliveryRatings, err := s.ratingRepo.GetDeliveryRatingByOrderId(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return deliveryRatings, nil

}

func (s *ratingSvcImpl) ListDeliveryRatingByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.DeliveryRating, error) {

	deliveryRatings, err := s.ratingRepo.ListDeliveryRatingByEaterId(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return deliveryRatings, nil
}

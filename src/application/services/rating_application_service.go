package service

import (
	"context"
	"fmt"
	"eater-service/src/application/dtos"
	ratingsvc "eater-service/src/domain/rating/services"
)

type RatingApplicationService interface {
	RateRestaurant(ctx context.Context, eaterId string, restaurantID string, rating int, comment string) (*dtos.GetRestaurantRatingResponse, error)
	UpdateRestaurantRating(ctx context.Context, restaurantRatingId string, rating int, comment string) (*dtos.GetRestaurantRatingResponse, error)
	ListRestaurantRatingEaterId(ctx context.Context, eaterId string) (*dtos.ListRestaurantRatingResponse, error)
	RateDelivery(ctx context.Context, eaterId string, orderId string, rating int, comment string) (*dtos.GetDeliveryRatingResponse, error)
	UpdateDeliveryRating(ctx context.Context, deliveryRatingId string, rating int, comment string) (*dtos.GetDeliveryRatingResponse, error)
	GetDeliveryRatingByOrderId(ctx context.Context, orderId string) (*dtos.GetDeliveryRatingResponse, error)
	ListDeliveryRatingByEaterId(ctx context.Context, eaterId string) (*dtos.ListDeliveryRatingResponse, error)
}

type ratingAppSvcImpl struct {
	ratingSvc ratingsvc.RatingService
}

func NewRatingApplicationService(
	ratingSvc ratingsvc.RatingService,
) RatingApplicationService {
	return &ratingAppSvcImpl{
		ratingSvc: ratingSvc,
	}
}

func (s *ratingAppSvcImpl) RateRestaurant(ctx context.Context, eaterId string, restaurantID string, rating int, comment string) (*dtos.GetRestaurantRatingResponse, error) {

	if eaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterId)
	}

	if restaurantID == "" {
		return nil, fmt.Errorf("Invalid or missing restaurantID: %s", restaurantID)
	}

	if rating == 0 {
		return nil, fmt.Errorf("Invalid or missing rating: %s", rating)
	}

	if comment == "" {
		return nil, fmt.Errorf("Invalid or missing comment: %s", comment)
	}

	restaurantRating, err := s.ratingSvc.RateRestaurant(ctx, eaterId, restaurantID, rating, comment)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetRestaurantRatingResponse(restaurantRating), nil
}

func (s *ratingAppSvcImpl) UpdateRestaurantRating(ctx context.Context, restaurantRatingId string, rating int, comment string) (*dtos.GetRestaurantRatingResponse, error) {

	if restaurantRatingId == "" {
		return nil, fmt.Errorf("Invalid or missing restaurantRatingId: %s", restaurantRatingId)
	}

	if rating == 0 {
		return nil, fmt.Errorf("Invalid or missing rating: %s", rating)
	}

	if comment == "" {
		return nil, fmt.Errorf("Invalid or missing comment: %s", comment)
	}

	restaurantRating, err := s.ratingSvc.UpdateRestaurantRating(ctx, restaurantRatingId, rating, comment)
	if err != nil {
		return nil, err
	}

	response := dtos.GetRestaurantRatingResponse{
		RestaurantRating: restaurantRating,
	}
	return &response, nil
}

func (s *ratingAppSvcImpl) ListRestaurantRatingEaterId(ctx context.Context, eaterID string, sort string, page int, pageSize int) (*dtos.ListRestaurantRatingResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eaterId: %s", eaterID)
	}

	restaurantRating, err := s.ratingSvc.ListRestaurantRatingEaterId(ctx, eaterID, sort , page , pageSize )
	if err != nil {
		return nil, err
	}

	return dtos.NewListRestaurantRatingResponse(restaurantRating), nil
}

func (s *ratingAppSvcImpl) RateDelivery(ctx context.Context, eaterId string, orderId string, rating int, comment string) (*dtos.GetDeliveryRatingResponse, error) {

	if eaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterId)
	}

	if orderId == "" {
		return nil, fmt.Errorf("Invalid or missing orderId: %s", orderId)
	}

	if rating == 0 {
		return nil, fmt.Errorf("Invalid or missing rating: %s", rating)
	}

	if comment == "" {
		return nil, fmt.Errorf("Invalid or missing comment: %s", comment)
	}

	deliveryRating, err := s.ratingSvc.RateDelivery(ctx, eaterId, orderId, rating, comment)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetDeliveryRatingResponse(deliveryRating), nil
}

func (s *ratingAppSvcImpl) UpdateDeliveryRating(ctx context.Context, deliveryRatingId string, rating int, comment string) (*dtos.GetDeliveryRatingResponse, error) {

	if deliveryRatingId == "" {
		return nil, fmt.Errorf("Invalid or missing deliveryRatingId: %s", deliveryRatingId)
	}

	if rating == 0 {
		return nil, fmt.Errorf("Invalid or missing rating: %s", rating)
	}

	if comment == "" {
		return nil, fmt.Errorf("Invalid or missing comment: %s", comment)
	}

	deliveryRating, err := s.ratingSvc.UpdateDeliveryRating(ctx, deliveryRatingId, rating, comment)
	if err != nil {
		return nil, err
	}

	response := dtos.GetDeliveryRatingResponse{
		DeliveryRating: deliveryRating,
	}
	return &response, nil
}

func (s *ratingAppSvcImpl) GetDeliveryRatingByOrderId(ctx context.Context, orderId string) (*dtos.GetDeliveryRatingResponse, error) {

	if orderId == "" {
		return nil, fmt.Errorf("Invalid or missing orderId: %s", orderId)
	}

	deliveryRating, err := s.ratingSvc.GetDeliveryRatingByOrderId(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetDeliveryRatingResponse(deliveryRating), nil
}

func (s *ratingAppSvcImpl) ListDeliveryRatingByEaterId(ctx context.Context, eaterId string, sort string, page, pageSize int) (*dtos.ListDeliveryRatingResponse, error) {

	if eaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eaterId: %s", eaterId)
	}

	deliveryRating, err := s.ratingSvc.ListDeliveryRatingByEaterId(ctx, eaterId,sort , page, pageSize )
	if err != nil {
		return nil, err
	}

	return dtos.NewListDeliveryRatingResponse(deliveryRating), nil
}

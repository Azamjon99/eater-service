package service

import (
	"context"
	"fmt"
	"eater-service/src/application/dtos"
	ratingsvc "eater-service/src/domain/rating/services"
	pb "eater-service/src/application/protos/eater"

)

type RatingApplicationService interface {
	RateRestaurant(ctx context.Context, request *pb.RateRestaurantRequest) (*pb.RateRestaurantResponse, error)
	UpdateRestaurantRating(ctx context.Context, request *pb.UpdateRestaurantRatingRequest) (*pb.UpdateRestaurantRatingResponse, error)
	ListRestaurantRatingByEaterId(ctx context.Context, request *pb.ListRestaurantRatingByEaterRequest) (*pb.ListRestaurantRatingByEaterResponse, error)
	RateDelivery(ctx context.Context, request *pb.RateDeliveryRequest) (*pb.RateDeliveryResponse, error)
	UpdateDeliveryRating(ctx context.Context, request *pb.UpdateDeliveryRatingRequest) (*pb.UpdateDeliveryRatingResponse, error)
	ListDeliveryRatingByEaterId(ctx context.Context, request *pb.ListDeliveryRatingByEaterRequest) (*pb.ListDeliveryRatingByEaterResponse, error)
	GetDeliveryRatingByOrder(ctx context.Context, request *pb.GetDeliveryRatingByOrderRequest) (*pb.GetDeliveryRatingByOrderResponse, error)
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

func (s *ratingAppSvcImpl) RateRestaurant(ctx context.Context, request *pb.RateRestaurantRequest) (*pb.RateRestaurantResponse, error) {
	restaurantRating, err := s.ratingSvc.RateRestaurant(ctx, request.EaterId, request.RestaurantId, request.Rating, request.Comment)
	if err != nil {
		return nil, err
	}

	pbRestaurantRating := &pb.RestaurantRating{
		Id:           restaurantRating.ID,
		EaterId:      restaurantRating.EaterID,
		RestaurantId: restaurantRating.RestaurantID,
		Rating:       restaurantRating.Rating, 
		Comment:      restaurantRating.Comment,
	}

	response := &pb.RateRestaurantResponse{
		Rating: pbRestaurantRating,
	}

	return response, nil
}
func (s *ratingAppSvcImpl) UpdateRestaurantRating(ctx context.Context, request *pb.UpdateRestaurantRatingRequest) (*pb.UpdateRestaurantRatingResponse, error) {
	restaurantRating, err := s.ratingSvc.UpdateRestaurantRating(ctx, request.RatingId, request.Rating, request.Comment)
	if err != nil {
		return nil, err
	}

	pbRestaurantRating := &pb.RestaurantRating{
		Id:           restaurantRating.ID,
		EaterId:      restaurantRating.EaterID,
		RestaurantId: restaurantRating.RestaurantID,
		Rating:       restaurantRating.Rating, 
		Comment:      restaurantRating.Comment,
	}

	response := &pb.UpdateRestaurantRatingResponse{
		Rating: pbRestaurantRating,
	}

	return response, nil
}


func (s *ratingAppSvcImpl) RateDelivery(ctx context.Context, request *pb.RateDeliveryRequest) (*pb.RateDeliveryResponse, error) {
	deliveryRating, err := s.ratingSvc.RateDelivery(ctx, request.EaterId, request.OrderId, request.Rating, request.Comment)
	if err != nil {
		return nil, err
	}
	pbRestaurantRating := &pb.DeliveryRating{
		Id:           deliveryRating.ID,
		EaterId:      deliveryRating.EaterID,
		Rating:       deliveryRating.Rating, 
		Comment:      deliveryRating.Comment,
	}

	response := &pb.RateDeliveryResponse{
		Rating: pbRestaurantRating,
	}

	return response, nil
}

func (s *ratingAppSvcImpl) UpdateDeliveryRating(ctx context.Context, request *pb.UpdateDeliveryRatingRequest) (*pb.UpdateDeliveryRatingResponse, error) {
	deliveryRating, err := s.ratingSvc.UpdateDeliveryRating(ctx, request.RatingId, request.Rating, request.Comment)
	if err != nil {
		return nil, err
	}

	pbRestaurantRating := &pb.DeliveryRating{
		Id:           deliveryRating.ID,
		EaterId:      deliveryRating.EaterID,
		Rating:       deliveryRating.Rating, 
		Comment:      deliveryRating.Comment,
	}

	response := &pb.UpdateDeliveryRatingResponse{
		Rating: pbRestaurantRating,
	}

	return response, nil
}

func (s *ratingAppSvcImpl) ListDeliveryRatingByEaterId(ctx context.Context, request *pb.ListDeliveryRatingByEaterRequest) (*pb.ListDeliveryRatingByEaterResponse, error) {
	deliveryRatings, err := s.ratingSvc.ListDeliveryRatingByEaterId(ctx, request.EaterId)
	if err != nil {
		return nil, err
	}

	pbRestaurantRating := &pb.ListRestaurantRatingByEaterResponse{
		ratings:           deliveryRatings.EaterID,

	}

	response := &pb.UpdateDeliveryRatingResponse{
		Rating: pbRestaurantRating,
	}

	return response, nil
}

func (s *ratingAppSvcImpl) GetDeliveryRatingByOrder(ctx context.Context, request *pb.GetDeliveryRatingByOrderRequest) (*pb.GetDeliveryRatingByOrderResponse, error) {
	deliveryRating, err := s.ratingSvc.GetDeliveryRatingByOrderId(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.GetDeliveryRatingByOrderResponse{Rating: deliveryRating}, nil
}






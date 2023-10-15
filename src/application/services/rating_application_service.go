package service

import (
	"context"

	"github.com/Azamjon99/eater-service/src/application/dtos"
	orderrated "github.com/Azamjon99/eater-service/src/application/integration_events/events/order_rated"
	restaurantrated "github.com/Azamjon99/eater-service/src/application/integration_events/events/restaurant_rated"
	restaurantratingupdated "github.com/Azamjon99/eater-service/src/application/integration_events/events/restaurant_rating_updated"
	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	ratingsvc "github.com/Azamjon99/eater-service/src/domain/rating/services"
	"github.com/Azamjon99/eater-service/src/infrastructure/pubsub"
	"go.uber.org/zap"
)

type RatingApplicationService interface {
	RateRestaurant(ctx context.Context, request *pb.RateRestaurantRequest) (*pb.RateRestaurantResponse, error)
	UpdateRestaurantRating(ctx context.Context, request *pb.UpdateRestaurantRatingRequest) (*pb.UpdateRestaurantRatingResponse, error)
	RateDelivery(ctx context.Context, request *pb.RateDeliveryRequest) (*pb.RateDeliveryResponse, error)
	UpdateDeliveryRating(ctx context.Context, request *pb.UpdateDeliveryRatingRequest) (*pb.UpdateDeliveryRatingResponse, error)
	ListDeliveryRatingByEaterId(ctx context.Context, request *pb.ListRestaurantRatingByEaterRequest) (*pb.ListRestaurantRatingByEaterResponse, error)
	GetDeliveryRatingByOrder(ctx context.Context, request *pb.GetDeliveryRatingByOrderRequest) (*pb.GetDeliveryRatingByOrderResponse, error)
}

type ratingAppSvcImpl struct {
	ratingSvc ratingsvc.RatingService
	producer  pubsub.Producer
	logger  *zap.Logger
}

func NewRatingApplicationService(
	ratingSvc ratingsvc.RatingService,
	producer pubsub.Producer,
	logger *zap.Logger,
) RatingApplicationService {
	return &ratingAppSvcImpl{
		ratingSvc: ratingSvc,
		producer: producer,
		logger: logger,
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

	event := restaurantrated.NewEvent(request.RestaurantId,restaurantRating.ID,request.Comment,int(request.Rating))
	if err := s.producer.PublishWithKey(event.Topic(),restaurantRating.ID,event.Payload(),nil);err!=nil{
		s.logger.Error("error pushing message",zap.Error(err),zap.String("topic",event.Topic()))
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

	event := restaurantratingupdated.NewEvent(request.RatingId,request.Comment,int(request.Rating))
	if err := s.producer.PublishWithKey(event.Topic(),request.RatingId,event.Payload(),nil);err!=nil{
		s.logger.Error("error pushing message",zap.Error(err),zap.String("topic",event.Topic()))
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
	event := orderrated.NewEvent(deliveryRating.ID,request.OrderId,request.Comment,int(request.Rating))
	if err := s.producer.PublishWithKey(event.Topic(),deliveryRating.ID,event.Payload(),nil);err!=nil{
		s.logger.Error("error pushing message",zap.Error(err),zap.String("topic",event.Topic()))
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
	event := orderrated.NewEvent(deliveryRating.ID,deliveryRating.OrderID,request.Comment,int(request.Rating))
	if err := s.producer.PublishWithKey(event.Topic(),deliveryRating.ID,event.Payload(),nil);err!=nil{
		s.logger.Error("error pushing message",zap.Error(err),zap.String("topic",event.Topic()))
	}
	response := &pb.UpdateDeliveryRatingResponse{
		Rating: pbRestaurantRating,
	}

	return response, nil
}

func (s *ratingAppSvcImpl) ListDeliveryRatingByEaterId(ctx context.Context, request *pb.ListRestaurantRatingByEaterRequest) (*pb.ListRestaurantRatingByEaterResponse, error) {
    deliveryRatings, err := s.ratingSvc.ListDeliveryRatingByEaterId(
        ctx,
        request.EaterId,

    )
    if err != nil {
        return nil, err
    }

	return &pb.ListRestaurantRatingByEaterResponse{
		Ratings: dtos.ToRestaurantRatingsPB(deliveryRatings),
	},nil
}

func (s *ratingAppSvcImpl) GetDeliveryRatingByOrder(ctx context.Context, request *pb.GetDeliveryRatingByOrderRequest) (*pb.GetDeliveryRatingByOrderResponse, error) {
	deliveryRating, err := s.ratingSvc.GetDeliveryRatingByOrderId(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.GetDeliveryRatingByOrderResponse{
		Rating: dtos.ToDeliveryRatingPB(deliveryRating),
	},nil}






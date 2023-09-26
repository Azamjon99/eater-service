package dtos

import (
	"time"

	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	"github.com/Azamjon99/eater-service/src/domain/rating/models"
)

func ToDeliveryRatingPB(rating *models.DeliveryRating) *pb.DeliveryRating {
	return &pb.DeliveryRating{
		Id:        rating.ID,
		EaterId:   rating.EaterID,
		OrderId:   rating.OrderID,
		Rating:    int32(rating.Rating),
		Comment:   rating.Comment,
		CreatedAt: rating.CreatedAt.Format(time.RFC3339),
		UpdatedAt: rating.UpdatedAt.Format(time.RFC3339),
	}
}

func ToDeliveryRatingsPB(ratings []*models.DeliveryRating) []*pb.DeliveryRating {
	ratingsArr := make([]*pb.DeliveryRating, len(ratings))
	for i, rating := range ratings {
		ratingsArr[i] = ToDeliveryRatingPB(rating)
	}
	return ratingsArr
}

func ToRestaurantRatingPB(rating *models.DeliveryRating) *pb.RestaurantRating {
	return &pb.RestaurantRating{
		Id:        rating.ID,
		EaterId:   rating.EaterID,
		Rating:    int32(rating.Rating),
		Comment:   rating.Comment,
		CreatedAt: rating.CreatedAt.Format(time.RFC3339),
		UpdatedAt: rating.UpdatedAt.Format(time.RFC3339),
	}
}

func ToRestaurantRatingsPB(ratings []*models.DeliveryRating) []*pb.RestaurantRating {
	ratingsArr := make([]*pb.RestaurantRating, len(ratings))
	for i, rating := range ratings {
		ratingsArr[i] = ToRestaurantRatingPB(rating)
	}
	return ratingsArr
}

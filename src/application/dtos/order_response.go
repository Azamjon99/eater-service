package dtos

import (

	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	"github.com/Azamjon99/eater-service/src/domain/order/models"
)



func ConvertToPbOrders(order *models.Order) *pb.Order {
	return &pb.Order{
		Id: order.ID,
		EaterId: order.EaterID,
	}
}


func ConvertToPbOrder(orders []*models.Order) []*pb.Order {
	ordersArr := make([]*pb.Order, len(orders))
	for i, order := range orders {
		ordersArr[i] = ConvertToPbOrders(order)
	}
	return ordersArr
}

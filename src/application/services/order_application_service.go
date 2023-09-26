package service

import (
	"context"
	"fmt"
	"eater-service/src/application/dtos"
	ordersvc "eater-service/src/domain/order/services"
	pb "eater-service/src/application/protos/address"

)

type OrderApplicationService interface {
	CreateOrder(ctx context.Context, request *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error)
	UpdateOrder(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	UpdateOrderByStatus(ctx context.Context, request *pb.UpdateOrderByStatusRequest) (*pb.UpdateOrderResponse, error)
	UpdateOrderPaymentByStatus(ctx context.Context, request *pb.UpdateOrderPaymentByStatusRequest) (*pb.UpdateOrderResponse, error)
	DeleteOrder(ctx context.Context, request *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
	GetOrderById(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	ListOrderByEaterId(ctx context.Context, request *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error)
}

type orderAppSvcImpl struct {
	orderSvc ordersvc.OrderService
} 

func NewOrderApplicationService(orderSvc ordersvc.OrderService) OrderApplicationService {
	return &orderAppSvcImpl{
		orderSvc: orderSvc,
	}
}

func (s *orderAppSvcImpl) CreateOrder(ctx context.Context, request *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	order, err := s.orderSvc.CreateOrder(ctx, request.EaterId, request.Cart)
	if err != nil {
		return nil, err
	}

	return &pb.PlaceOrderResponse{Order: order}, nil
}

func (s *orderAppSvcImpl) UpdateOrder(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order, err := s.orderSvc.UpdateOrder(ctx, request.Order)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderResponse{Order: order}, nil
}

func (s *orderAppSvcImpl) UpdateOrderByStatus(ctx context.Context, request *pb.UpdateOrderByStatusRequest) (*pb.UpdateOrderResponse, error) {
	order, err := s.orderSvc.UpdateOrderByStatus(ctx, request.Order)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderResponse{Order: order}, nil
}

func (s *orderAppSvcImpl) UpdateOrderPaymentByStatus(ctx context.Context, request *pb.UpdateOrderPaymentByStatusRequest) (*pb.UpdateOrderResponse, error) {
	order, err := s.orderSvc.UpdateOrderPaymentByStatus(ctx, request.Order)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateOrderResponse{Order: order}, nil
}

func (s *orderAppSvcImpl) DeleteOrder(ctx context.Context, request *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	_, err := s.orderSvc.DeleteOrder(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteOrderResponse{}, nil
}

func (s *orderAppSvcImpl) GetOrderById(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := s.orderSvc.GetOrderById(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderResponse{Order: order}, nil
}

func (s *orderAppSvcImpl) ListOrderByEaterId(ctx context.Context, request *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error) {
	orders, err := s.orderSvc.ListOrderByEaterId(ctx, request.EaterId, request.Sort, int(request.Page), int(request.PageSize))
	if err != nil {
		return nil, err
	}

	return &pb.ListOrderByEaterResponse{Orders: orders}, nil
}
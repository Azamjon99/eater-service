package service

import (
	"context"
	"fmt"
	"eater-service/src/application/dtos"
	ordersvc "eater-service/src/domain/order/services"
)

type OrderApplicationService interface {
	CreateOrder(ctx context.Context, eaterId string, instruction string, restaurantID string) (*dtos.GetOrderResponse, error)
	UpdateOrder(ctx context.Context, orderId string, instruction string, restaurantID string) (*dtos.GetOrderResponse, error)
	UpdateOrderByStatus(ctx context.Context, orderId string, status string) (*dtos.GetOrderResponse, error)
	UpdateOrderPaymentByStatus(ctx context.Context, orderId string, paymentStatus string) (*dtos.GetOrderResponse, error)
	DeleteOrder(ctx context.Context, orderId string) (*dtos.GetOrderResponse, error)
	GetOrderById(ctx context.Context, orderId string) (*dtos.GetOrderResponse, error)
	ListOrderByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListOrderResponse, error)
}

type orderAppSvcImpl struct {
	orderSvc ordersvc.OrderService
}

func NewOrderApplicationService(orderSvc ordersvc.OrderService) OrderApplicationService {
	return &orderAppSvcImpl{
		orderSvc: orderSvc,
	}
}

func (s *orderAppSvcImpl) CreateOrder(ctx context.Context, eaterId string, instruction string, restaurantID string) (*dtos.GetOrderResponse, error) {

	if eaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterId)
	}
	if instruction == "" {
		return nil, fmt.Errorf("Invalid or missing instruction: %s", instruction)
	}
	if restaurantID == "" {
		return nil, fmt.Errorf("Invalid or missing restaurantID: %s", restaurantID)
	}

	order, err := s.orderSvc.CreateOrder(ctx, eaterId, instruction, restaurantID)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetOrderResponse(order), nil
}

func (s *orderAppSvcImpl) UpdateOrder(ctx context.Context, orderId string, instruction string, restaurantID string) (*dtos.GetOrderResponse, error) {
	if orderId == "" {
		return nil, fmt.Errorf("Invalid or missing orderId: %s", orderId)
	}
	if instruction == "" {
		return nil, fmt.Errorf("Invalid or missing instruction: %s", instruction)
	}
	if restaurantID == "" {
		return nil, fmt.Errorf("Invalid or missing restaurantID: %s", restaurantID)
	}

	order, err := s.orderSvc.UpdateOrder(ctx, orderId, instruction, restaurantID)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetOrderResponse(order), nil
}

func (s *orderAppSvcImpl) UpdateOrderByStatus(ctx context.Context, orderId string, status string) (*dtos.GetOrderResponse, error) {
	if orderId == "" {
		return nil, fmt.Errorf("Invalid or missing orderId: %s", orderId)
	}
	if status == "" {
		return nil, fmt.Errorf("Invalid or missing status: %s", status)
	}

	order, err := s.orderSvc.UpdateOrderByStatus(ctx, orderId, status)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetOrderResponse(order), nil
}

func (s *orderAppSvcImpl) UpdateOrderPaymentByStatus(ctx context.Context, orderId string, paymentStatus string) (*dtos.GetOrderResponse, error) {
	if orderId == "" {
		return nil, fmt.Errorf("Invalid or missing orderId: %s", orderId)
	}
	if paymentStatus == "" {
		return nil, fmt.Errorf("Invalid or missing paymentStatus: %s", paymentStatus)
	}

	order, err := s.orderSvc.UpdateOrderPaymentByStatus(ctx, orderId, paymentStatus)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetOrderResponse(order), nil
}

func (s *orderAppSvcImpl) DeleteOrder(ctx context.Context, orderId string) (*dtos.GetOrderResponse, error) {
	if orderId == "" {
		return nil, fmt.Errorf("Invalid or missing orderId: %s", orderId)
	}

	order, err := s.orderSvc.DeleteOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetOrderResponse(order), nil
}

func (s *orderAppSvcImpl) GetOrderById(ctx context.Context, orderId string) (*dtos.GetOrderResponse, error) {

	if orderId == "" {
		return nil, fmt.Errorf("Invalid or missing orderId: %s", orderId)
	}

	order, err := s.orderSvc.GetOrderById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetOrderResponse(order), nil
}

func (s *orderAppSvcImpl) ListOrderByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListOrderResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eaterID: %s", eaterID)
	}

	order, err := s.orderSvc.ListOrderByEaterId(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return dtos.NewListOrderResponse(order), nil
}

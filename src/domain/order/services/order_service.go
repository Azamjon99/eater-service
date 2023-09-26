package services

import (
	"context"
	"github.com/Azamjon99/eater-service/src/domain/order/models"
	"github.com/Azamjon99/eater-service/src/domain/order/repositories"
	"github.com/Azamjon99/eater-service/src/infrastructure/rand"
	"time"
)

type OrderService interface {
	CreateOrder(ctx context.Context, eaterId string, instruction string, restaurantID string) (*models.Order, error)
	UpdateOrder(ctx context.Context, orderId string, instruction string, restaurantID string) (*models.Order, error)
	UpdateOrderByStatus(ctx context.Context, orderId string, status string) (*models.Order, error)
	UpdateOrderPaymentByStatus(ctx context.Context, orderId string, paymentStatus string) (*models.Order, error)
	DeleteOrder(ctx context.Context, orderId string) (*models.Order, error)
	GetOrderById(ctx context.Context, orderId string) (*models.Order, error)
	ListOrderByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Order, error)
}

type orderSvcImpl struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService(orderRepo repositories.OrderRepository) *orderSvcImpl {
	return &orderSvcImpl{
		orderRepo: orderRepo,
	}
}

func (s *orderSvcImpl) CreateOrder(ctx context.Context, eaterId string, instruction string, restaurantID string) (*models.Order, error) {

	order := &models.Order{
		ID:           rand.UUID(),
		EaterID:      eaterId,
		Instruction:  instruction,
		RestaurantID: restaurantID,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	err := s.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil

}

func (s *orderSvcImpl) UpdateOrder(ctx context.Context, orderId string, instruction string, restaurantID string) (*models.Order, error) {

	order := &models.Order{
		ID:           orderId,
		Instruction:  instruction,
		RestaurantID: restaurantID,
		UpdatedAt:    time.Now().UTC(),
	}

	
	 err := s.orderRepo.UpdateOrder(ctx, order);
	if err != nil {
		return nil, err
	}

	return order, nil

}

func (s *orderSvcImpl) UpdateOrderByStatus(ctx context.Context, orderId string, status string) (*models.Order, error) {

	order := &models.Order{
		EaterID:   orderId,
		Status:    status,
		UpdatedAt: time.Now().UTC(),
	}

	 err := s.orderRepo.UpdateOrderByStatus(ctx, orderId, status); 

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderSvcImpl) UpdateOrderPaymentByStatus(ctx context.Context, orderId string, paymentStatus string) (*models.Order, error) {

	order := &models.Order{
		EaterID:       orderId,
		PaymentStatus: paymentStatus,
		UpdatedAt:     time.Now().UTC(),
	}

	order, err := s.UpdateOrderPaymentByStatus(ctx, orderId, paymentStatus); 
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderSvcImpl) DeleteOrder(ctx context.Context, orderId string) (*models.Order, error) {

	 err :=   s.orderRepo.DeleteOrder(ctx, orderId);

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *orderSvcImpl) GetOrderById(ctx context.Context, orderId string) (*models.Order, error) {

	order, err := s.orderRepo.GetOrderById(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return order, nil

}

func (s *orderSvcImpl) ListOrderByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Order, error) {

	orders, err := s.orderRepo.ListOrderByEaterId(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return orders, nil

}

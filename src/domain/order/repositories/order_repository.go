package repositories

import (
	"context"

	"eater-service/src/domain/order/models"
)

type OrderRepository interface {
	WithTx(ctx context.Context, f func(r OrderRepository) error) error
	CreateOrder(ctx context.Context, orders *models.Order) error
	UpdateOrder(ctx context.Context, orders *models.Order) error
	UpdateOrderByStatus(ctx context.Context, orderId string, status string) error
	UpdateOrderPaymentByStatus(ctx context.Context, orderId string, paymentStatus string) error
	DeleteOrder(ctx context.Context, orderId string) error
	GetOrderById(ctx context.Context, orderId string) (*models.Order, error)
	ListOrderByEaterId(ctx context.Context, eaterId string, sort string, page, pageSize int) ([]*models.Order, error)
}

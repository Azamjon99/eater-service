package order

import (
	"context"

	"github.com/Azamjon99/eater-service/src/domain/order/models"
	"github.com/Azamjon99/eater-service/src/domain/order/repositories"
	"github.com/Azamjon99/eater-service/src/infrastructure/repositories/utils"
	"gorm.io/gorm"
)

const (
	tableOrder = "eater.orders"
)

type orderRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderRepoImpl{
		db: db,
	}
}

func (r *orderRepoImpl) WithTx(ctx context.Context, callback func(r repositories.OrderRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := orderRepoImpl{
			db: tx,
		}
		if err := callback(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *orderRepoImpl) CreateOrder(ctx context.Context, orders *models.Order) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Create(orders)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) UpdateOrder(ctx context.Context, orders *models.Order) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Save(orders)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) UpdateOrderByStatus(ctx context.Context, orderId string, status string) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Where("id", orderId).Update("status = ", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) UpdateOrderPaymentByStatus(ctx context.Context, orderId string, paymentStatus string) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Where("id", orderId).Update("payment_status = ", paymentStatus)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) DeleteOrder(ctx context.Context, orderId string) error {
	var order *models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).Delete(&order, "id = ?", orderId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderRepoImpl) GetOrderById(ctx context.Context, orderId string) (*models.Order, error) {
	var order *models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).First(&order, "id = ?", orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (r *orderRepoImpl) ListOrderByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Order, error) {
	var orders []*models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).Where("eater_id = ?", eaterID)
	result.Scopes(utils.SortByCreatedAt(page, pageSize), utils.Sort(sort)).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

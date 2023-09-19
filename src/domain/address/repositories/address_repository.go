package repositories

import (
	"context"
	"eater-service/src/domain/address/models"
)

type AddressRepository interface {
	WithTx(ctx context.Context, f func(r AddressRepository) error) error
	CreateAddress(ctx context.Context, address *models.Address) error
	UpdateAddress(ctx context.Context, address *models.Address) error
	DeleteAddress(ctx context.Context, addressId string) error
	GetAddressById(ctx context.Context, addressId string) (*models.Address, error)
	ListAddressByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error)
}
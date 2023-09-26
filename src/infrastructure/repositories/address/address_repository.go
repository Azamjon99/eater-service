package address

import (
	"context"

	"github.com/Azamjon99/eater-service/src/domain/address/models"
	"github.com/Azamjon99/eater-service/src/domain/address/repositories"
	"github.com/Azamjon99/eater-service/src/infrastructure/repositories/utils"
	"gorm.io/gorm"
)

const (
	tableAddress = "address.address"
)

type addressRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.AddressRepository {
	return &addressRepoImpl{
		db: db,
	}
}

func (r *addressRepoImpl) CreateAddress(ctx context.Context, address *models.Address) error {
	result := r.db.WithContext(ctx).Table(tableAddress).Create(address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *addressRepoImpl) UpdateAddress(ctx context.Context, address *models.Address) error {
	result := r.db.WithContext(ctx).Table(tableAddress).Save(address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *addressRepoImpl) DeleteAddress(ctx context.Context, addressId string) error {
	var address models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Delete(&address, "id = ?", addressId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *addressRepoImpl) GetAddressById(ctx context.Context, addressId string) (*models.Address, error) {
	var address *models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).First(&address, "id = ?", addressId)
	if result.Error != nil {
		return nil, result.Error
	}

	return address, nil
}

func (r *addressRepoImpl) ListAddressByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error) {
	var addresses []*models.Address
	result := r.db.WithContext(ctx).Table(tableAddress).Where("eater_id = ?", eaterID)
	result.Scopes(utils.SortByCreatedAt(page, pageSize), utils.Sort(sort)).Find(&addresses)

	if result.Error != nil {
		return nil, result.Error
	}
	return addresses, nil
}

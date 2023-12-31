package services

import (
	"context"
	"github.com/Azamjon99/eater-service/src/domain/address/models"
	"github.com/Azamjon99/eater-service/src/domain/address/repositories"
	"github.com/Azamjon99/eater-service/src/infrastructure/rand"
	"time"
)


type AddressService interface {
	SaveAddress(ctx context.Context, eaterId string, name string, long, lat float64) (*models.Address, error)
	UpdateAddress(ctx context.Context, addressId string, name string, long, lat float64) (*models.Address, error)
	DeleteAddress(ctx context.Context, addressId string) (*models.Address, error)
	GetAddressById(ctx context.Context, addressId string) (*models.Address, error)
	ListAddressByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error)
}
type addressSvcImpl struct {
	addressRepo repositories.AddressRepository
}

func NewAddressService(addressRepo repositories.AddressRepository) AddressService {
	return &addressSvcImpl{
		addressRepo: addressRepo,
	}
}

func (s *addressSvcImpl) SaveAddress(ctx context.Context, eaterId string, name string, long, lat float64) (*models.Address, error) {

	location := &models.Location{
		Longitude: long,
		Latitude:  lat,
	}

	address := &models.Address{
		ID:        rand.UUID(),
		EaterID:   eaterId,
		Name:      name,
		Location:  location,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.addressRepo.SaveAddress(ctx,address)

	return address,err

}

func (s *addressSvcImpl) UpdateAddress(ctx context.Context, addressId string, name string, long, lat float64) (*models.Address, error) {

	location := &models.Location{
		Longitude: long,
		Latitude:  lat,
	}

	address := &models.Address{
		ID:        addressId,
		Name:      name,
		Location:  location,
		UpdatedAt: time.Now().UTC(),
	}

		
	err := s.addressRepo.UpdateAddress(ctx,address)

	return address,err
}

func (s *addressSvcImpl) DeleteAddress(ctx context.Context, addressId string) (*models.Address, error) {

	

	return s.DeleteAddress(ctx, addressId)
}

func (s *addressSvcImpl) GetAddressById(ctx context.Context, addressId string) (*models.Address, error) {

	address, err := s.addressRepo.GetAddressById(ctx, addressId)
	if err != nil {
		return nil, err
	}

	return address, nil

}

func (s *addressSvcImpl) ListAddressByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Address, error) {

	address, err := s.addressRepo.ListAddressByEaterId(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return address, nil
}

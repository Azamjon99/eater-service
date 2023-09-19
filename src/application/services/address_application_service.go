package service

import (
	"context"
	"fmt"
	"eater-service/src/application/dtos"
	addresssvc "eater-service/src/domain/address/services"
)

type AddressApplicationService interface {
	CreateAddress(ctx context.Context, eaterId string, name string, long, lat float64) (*dtos.GetAddressResponse, error)
	UpdateAddress(ctx context.Context, addressId string, name string, long, lat float64) (*dtos.UpdateAddressResponse, error)
	DeleteAddress(ctx context.Context, addressId string) (*dtos.GetAddressResponse, error)
	GetAddressById(ctx context.Context, addressId string) (*dtos.GetAddressResponse, error)
	ListAddressByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListAddressResponse, error)
}

type addressAppSvcImpl struct {
	addressSvc addresssvc.AddressService
}

func NewAddressApplicationService(
	addressSvc addresssvc.AddressService,
) AddressApplicationService {
	return &addressAppSvcImpl{
		addressSvc: addressSvc,
	}
}

func (s *addressAppSvcImpl) CreateAddress(ctx context.Context, eaterId string, name string, long, lat float64) (*dtos.GetAddressResponse, error) {

	if eaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterId)
	}

	if name == "" {
		return nil, fmt.Errorf("Invalid or missing name: %s", name)
	}

	if long == 0 {
		return nil, fmt.Errorf("Invalid or missing long: %s", long)
	}

	if lat == 0 {
		return nil, fmt.Errorf("Invalid or missing lat: %s", lat)
	}

	address, err := s.addressSvc.CreateAddress(ctx, eaterId, name, long, lat)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetAddressResponse(address), nil
}

func (s *addressAppSvcImpl) UpdateAddress(ctx context.Context, addressId string, name string, long, lat float64) (*dtos.UpdateAddressResponse, error) {

	if addressId == "" {
		return nil, fmt.Errorf("Invalid or missing address_id: %s", addressId)
	}

	if name == "" {
		return nil, fmt.Errorf("Invalid or missing name: %s", name)
	}

	if long == 0 {
		return nil, fmt.Errorf("Invalid or missing long: %s", long)
	}

	if lat == 0 {
		return nil, fmt.Errorf("Invalid or missing lat: %s", lat)
	}

	address, err := s.addressSvc.UpdateAddress(ctx, addressId, name, long, lat)
	if err != nil {
		return nil, err
	}

	response := dtos.UpdateAddressResponse{
		Address: address,
	}
	return &response, nil
}

func (s *addressAppSvcImpl) DeleteAddress(ctx context.Context, addressId string) (*dtos.GetAddressResponse, error) {

	if addressId == "" {
		return nil, fmt.Errorf("Invalid or missing addressId: %s", addressId)
	}

	address, err := s.addressSvc.DeleteAddress(ctx, addressId)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetAddressResponse(address), nil
}

func (s *addressAppSvcImpl) GetAddressById(ctx context.Context, addressId string) (*dtos.GetAddressResponse, error) {

	if addressId == "" {
		return nil, fmt.Errorf("Invalid or missing addressId: %s", addressId)
	}

	address, err := s.addressSvc.GetAddressById(ctx, addressId)
	if err != nil {
		return nil, err
	}

	return dtos.NewGetAddressResponse(address), nil
}

func (s *addressAppSvcImpl) ListAddressByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) (*dtos.ListAddressResponse, error) {

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing addeaterIDressId: %s", eaterID)
	}
	

	addresses, err := s.addressSvc.ListAddressByEaterId(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return dtos.NewListAddressResponse(addresses), nil
}

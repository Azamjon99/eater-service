package service

import (
	"context"
	"fmt"
	"eater-service/src/application/dtos"
	addresssvc "eater-service/src/domain/address/services"
	pb "eater-service/src/application/protos/address"

)

type AddressApplicationService interface {
	CreateAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error)
	UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error)
	DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error)
	GetAddressById(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error)
	ListAddressByEaterId(ctx context.Context, req *pb.ListAddressByEaterRequest) (*pb.ListAddressByEaterResponse, error)
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

func (s *addressAppSvcImpl) CreateAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	if req.EaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", req.EaterId)
	}

	if req.Name == "" {
		return nil, fmt.Errorf("Invalid or missing name: %s", req.Name)
	}

	if req.Longitude == 0 {
		return nil, fmt.Errorf("Invalid or missing longitude: %f", req.Longitude)
	}

	if req.Latitude == 0 {
		return nil, fmt.Errorf("Invalid or missing latitude: %f", req.Latitude)
	}

	address, err := s.addressSvc.CreateAddress(ctx, req.EaterId, req.Name, req.Longitude, req.Latitude)
	if err != nil {
		return nil, err
	}

	return &pb.AddAddressResponse{
		Address: address,
	}, nil
}

func (s *addressAppSvcImpl) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error) {
	if req.AddressId == "" {
		return nil, fmt.Errorf("Invalid or missing address_id: %s", req.AddressId)
	}

	if req.Name == "" {
		return nil, fmt.Errorf("Invalid or missing name: %s", req.Name)
	}

	if req.Longitude == 0 {
		return nil, fmt.Errorf("Invalid or missing longitude: %f", req.Longitude)
	}

	if req.Latitude == 0 {
		return nil, fmt.Errorf("Invalid or missing latitude: %f", req.Latitude)
	}

	address, err := s.addressSvc.UpdateAddress(ctx, req.AddressId, req.Name, req.Longitude, req.Latitude)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAddressResponse{
		Address: address,
	}, nil
}

func (s *addressAppSvcImpl) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressResponse, error) {
	if req.AddressId == "" {
		return nil, fmt.Errorf("Invalid or missing address_id: %s", req.AddressId)
	}

	_, err := s.addressSvc.DeleteAddress(ctx, req.AddressId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAddressResponse{}, nil
}

func (s *addressAppSvcImpl) GetAddressById(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressResponse, error) {
	if req.AddressId == "" {
		return nil, fmt.Errorf("Invalid or missing address_id: %s", req.AddressId)
	}

	address, err := s.addressSvc.GetAddressById(ctx, req.AddressId)
	if err != nil {
		return nil, err
	}

	return &pb.GetAddressResponse{
		Address: address,
	}, nil
}

func (s *addressAppSvcImpl) ListAddressByEaterId(ctx context.Context, req *pb.ListAddressByEaterRequest) (*pb.ListAddressByEaterResponse, error) {
	if req.EaterId == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", req.EaterId)
	}

	addresses, err := s.addressSvc.ListAddressByEaterId(ctx, req.EaterId)
	if err != nil {
		return nil, err
	}

	return &pb.ListAddressByEaterResponse{
		Addresses: addresses,
	}, nil
}
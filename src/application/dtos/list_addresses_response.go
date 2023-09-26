package dtos

import "github.com/Azamjon99/eater-service/src/domain/address/models"

type GetAddressListResponse struct {
	Addresses []*models.Address `json:"addresses"`
}

func NewGetAddressListResponse(addresses []*models.Address) *GetAddressListResponse {
	return &GetAddressListResponse{
		Addresses: addresses,
	}
}
package dtos

import "github.com/Azamjon99/eater-service/src/domain/address/models"

type ListAddressResponse struct {
	Address []*models.Address `json:"address"`
}

func NewListAddressResponse(address []*models.Address) *ListAddressResponse {
	return &ListAddressResponse{
		Address: address,
	}
}

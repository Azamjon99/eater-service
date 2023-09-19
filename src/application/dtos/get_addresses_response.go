package dtos

import "eater-service/src/domain/address/models"


type GetAddressResponse struct {
	Address *models.Address `json:"address"`
}

func NewGetAddressResponse(address *models.Address) *GetAddressResponse {
	return &GetAddressResponse{
		Address: address,
	}
}

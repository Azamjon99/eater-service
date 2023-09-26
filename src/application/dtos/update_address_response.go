package dtos

import "github.com/Azamjon99/eater-service/src/domain/address/models"

type UpdateAddressResponse struct {
	Address *models.Address `json:"address"`
}

func NewUpdateAddressResponse(address *models.Address) *UpdateAddressResponse {
	return &UpdateAddressResponse{
		Address: address,
	}
}

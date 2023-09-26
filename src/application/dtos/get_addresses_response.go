package dtos

import (
	"time"

	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	"github.com/Azamjon99/eater-service/src/domain/address/models"
)

type GetAddressResponse struct {
	Address *models.Address `json:"address"`
}

func NewGetAddressResponse(address *models.Address) *pb.Address {
	return &pb.Address{
		Id: address.ID,
		EaterId: address.EaterID,
		Name: address.Name,
		Location: &pb.Location{
			Longitude: address.Location.Longitude,
			Latitude: address.Location.Latitude,
		},
		CreatedAt: address.CreatedAt.Format(time.RFC3339),
		UpdatedAt: address.UpdatedAt.Format(time.RFC3339),
	}
}

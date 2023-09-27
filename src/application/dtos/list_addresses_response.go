package dtos

import ("github.com/Azamjon99/eater-service/src/domain/address/models"
pb "github.com/Azamjon99/eater-service/src/application/protos/eater")



func ConvertToPbAddresses(Address *models.Address) *pb.Address {
	return &pb.Address{
		Id: Address.ID,
		EaterId: Address.EaterID,
	}
}


func ConvertToPbAddress(addresses []*models.Address) []*pb.Address {
	addressArr := make([]*pb.Address, len(addresses))
    for i, address := range addresses {
        addressArr[i] = ConvertToPbAddresses(address)
    }
    return addressArr
}

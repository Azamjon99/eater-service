package dtos

import (
	"time"

	pb "eater-service/src/application/protos/eater"
	"eater-service/src/domain/eater/models"
)
func EaterProfileResponse(profile *models.EaterProfile) *pb.EaterProfile{
	return &pb.EaterProfile{
		EaterId: profile.EaterID,
		PhoneNumber: profile.PhoneNumber,
		Name: profile.Name,
		IsPhoneNumberConfirmed: profile.IsPhoneNumberConfirmed,
		CreatedAt: profile.CreatedAt.Format(time.RFC3339),
		UpdatedAt: profile.UpdatedAt.Format(time.RFC3339),
	}
}
package dtos

import "eater-service/src/domain/eater/models"



type UpdateEaterProfileResponse struct {
	Profile *models.EaterProfile `json:"profile"`
}

func NewUpdateEaterProfileResponse(profile *models.EaterProfile) *UpdateEaterProfileResponse {
	return &UpdateEaterProfileResponse{
		Profile: profile,
	}
}

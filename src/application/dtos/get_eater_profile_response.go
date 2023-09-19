package dtos

import "eater-service/src/domain/eater/models"


type GetEaterProfileResponse struct {
	Profile *models.EaterProfile `json:"profile"`
}

func NewGetEaterProfileResponse(profile *models.EaterProfile) *GetEaterProfileResponse {
	return &GetEaterProfileResponse{
		Profile: profile,
	}
}

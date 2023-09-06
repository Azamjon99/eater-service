package services

import (
	"context"
	"eater-service/src/domain/eater/models"
	"microservices-course/lesson8/clean/eaterservice/infrastructure/sms"
	"microservices-course/lesson8/layered/eaterservice/repositories"
)



type EaterService interface {
	SignupEater(ctx context.Context, phoneNumber string) (string, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*models.EaterProfile, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*models.EaterProfile, error)
	}

type EaterSvcImpl struct {
	eaterRepo repositories.EaterRepository
	smsClient sms.Client
	logger    *zap.logger      
}

func newEaterService (
	eaterRepo repositories.EaterRepository, //dependency injection
	smsClient  sms.Client,
	logger  *zap.Logger,
) EaterService {
	return &EaterSvcImpl{
		eaterRepo :eaterRepo,
		smsClient :smsClient,
		logger:logger,
	}
}


func (s *EaterSvcImpl) SignupEater(ctx context.Context, phoneNumber string) (string, error){
	
	phoneNumberExists := true

	eater, err := s.eaterRepo.GetEaterByPhoneNumber(ctx, phoneNumber)

	if err != nil {
		phoneNumberExists = false
	} 

	if phoneNumberExists {
		return s.handExiststingEater(ctx, eater.ID)
	}
	return s.handleNewEater(ctx, phoneNumber)
}
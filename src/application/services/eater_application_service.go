package service

import (
	"context"
	"eater-service/src/application/dtos"
	pb "eater-service/src/application/protos/eater"
	eatersvc "eater-service/src/domain/eater/services"
	"eater-service/src/infrastructure/jwt"
	"eater-service/src/infrastructure/validator"
	"errors"
	"fmt"
)
 
type EaterApplicationService interface {
	SignupEater(ctx context.Context, req *pb.SignupEaterRequest) (*pb.SignupEaterResponse, error)
	ConfirmSMSCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error)
	UpdateEaterProfile(ctx context.Context, req *pb.UpdateEaterProfileRequest) (*pb.UpdateEaterProfileResponse, error)
	GetEaterProfile(ctx context.Context, req *pb.GetEaterProfileRequest) (*pb.GetEaterProfileResponse, error)
}

type eaterAppSvcImpl struct {
	eaterSvc eatersvc.EaterService
	tokenSvc jwt.Service
}

func NewEaterApplicationService(
	eaterSvc eatersvc.EaterService,
	tokenSvc jwt.Service,
) EaterApplicationService {
	return &eaterAppSvcImpl{
		eaterSvc: eaterSvc,
		tokenSvc: tokenSvc,
	}
}

func (s *eaterAppSvcImpl) SignupEater(ctx context.Context, req *pb.SignupEaterRequest) (*pb.SignupEaterResponse, error) {
	if !validator.ValidateUzPhoneNumber(req.PhoneNumber) {
		return nil, errors.New("invalid phone number")
	}
	eaterId, err := s.eaterSvc.SignupEater(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return &pb.SignupEaterResponse{
		EaterId: eaterId,
	}, nil
}

func (s *eaterAppSvcImpl) ConfirmSMSCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error) {
	eaterID := req.EaterId
	smsCode := req.SmsCode

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterID)
	}

	if smsCode == "" {
		return nil, fmt.Errorf("Invalid or missing sms_code: %s", smsCode)
	}
	profile, err := s.eaterSvc.ConfirmSMSCode(ctx, eaterID, smsCode)
	if err != nil {
		return nil, err
	}

	token, err := s.tokenSvc.CreateToken(ctx, profile.EaterID)
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmSmsCodeResponse{
		Profile: dtos.EaterProfileResponse(profile),

		Token: token,
	}, nil
}

func (s *eaterAppSvcImpl) UpdateEaterProfile(ctx context.Context, req *pb.UpdateEaterProfileRequest) (*pb.UpdateEaterProfileResponse, error) {
	eaterID := req.EaterId
	name := req.Name
	imageUrl := req.ImageUrl

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eater_id: %s", eaterID)
	}

	if name == "" {
		return nil, fmt.Errorf("Invalid or missing name: %s", name)
	}

	profile, err := s.eaterSvc.UpdateEaterProfile(ctx, eaterID, name, imageUrl)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateEaterProfileResponse{
		Profile: dtos.EaterProfileResponse(profile),

	}, nil
}

func (s *eaterAppSvcImpl) GetEaterProfile(ctx context.Context, req *pb.GetEaterProfileRequest) (*pb.GetEaterProfileResponse, error) {
	eaterID := req.EaterId

	if eaterID == "" {
		return nil, fmt.Errorf("Invalid or missing eater id")
	}
	profile, err := s.eaterSvc.GetEaterProfile(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	return &pb.GetEaterProfileResponse{
		Profile: dtos.EaterProfileResponse(profile),

	}, nil
}
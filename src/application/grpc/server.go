package grpc

import (
	"context"
	pb "eater-service/src/application/protos/eater"
	service "eater-service/src/application/services"
)
type Server struct{
	eaterApp service.EaterApplicationService
	addressApp service.AddressApplicationService
}


func NewServer(
	eaterApp service.EaterApplicationService,
	addressApp service.AddressApplicationService,
) *Server{
	return &Server{
		eaterApp: eaterApp,
	
	}
}

func (s *Server) SignupEater(ctx context.Context, r *pb.SignupEaterRequest)(*pb.SignupEaterResponse,error){
	return s.eaterApp.SignupEater(ctx,r)
}

func (s *Server) ConfirmSMSCode(ctx context.Context, r *pb.ConfirmSmsCodeRequest)(*pb.ConfirmSmsCodeResponse,error){
	return s.eaterApp.ConfirmSMSCode(ctx,r)
}
func (s *Server) UpdateEaterProfile(ctx context.Context, r *pb.UpdateEaterProfileRequest)(*pb.UpdateEaterProfileResponse,error){
	return s.eaterApp.UpdateEaterProfile(ctx,r)
}
func (s *Server) GetEaterProfile(ctx context.Context, r *pb.GetEaterProfileRequest)(*pb.GetEaterProfileResponse,error){
	return s.eaterApp.GetEaterProfile(ctx,r)
}

func (s *Server) CreateAddress(ctx context.Context, r *pb.AddAddressRequest)(*pb.AddAddressResponse, error){
	return s.addressApp.CreateAddress(ctx,r)
}

func (s *Server) UpdateAddress(ctx context.Context, r *pb.UpdateAddressRequest)(*pb.UpdateAddressResponse, error){
	return s.addressApp.UpdateAddress(ctx,r)
}

func (s *Server) DeleteAddress(ctx context.Context, r *pb.GetAddressRequest)(*pb.GetAddressResponse, error){
	return s.addressApp.DeleteAddress(ctx,r)
}

func (s *Server) ListAddressByEaterId(ctx context.Context, r *pb.DeleteAddressRequest)(*pb.DeleteAddressResponse, error){
	return s.addressApp.ListAddressByEaterId(ctx,r)
}
package grpc

import (
	"context"
	pb "eater-service/src/application/protos/eater"
	service "eater-service/src/application/services"
)
type Server struct{
	eaterApp service.EaterApplicationService
}


func NewServer(
	eaterApp service.EaterApplicationService,
) *Server{
	return &Server{
		eaterApp: eaterApp,
	
	}
}

func (s *Server) SignupEater(ctx context.Context, r *pb.SignupEaterRequest)(*pb.SignupEaterResponse,error){
	return s.eaterApp.SignupEater(ctx,r)
}

func (s *Server) ConfirmSMSCode(ctx context.Context, r *pb.SignupEaterRequest)(*pb.SignupEaterResponse,error){
	return s.eaterApp.SignupEater(ctx,r)
}


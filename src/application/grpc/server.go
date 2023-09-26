package grpc

import (
	"context"
	pb "eater-service/src/application/protos/eater"
	service "eater-service/src/application/services"
)
type Server struct{
	eaterApp service.EaterApplicationService
	addressApp service.AddressApplicationService
	restaurantRatingApp service.RatingApplicationService
}


func NewServer(
	eaterApp service.EaterApplicationService,
	addressApp service.AddressApplicationService,
	restaurantRatingApp service.RatingApplicationService,
) *Server{
	return &Server{
		eaterApp: eaterApp,
		addressApp: addressApp,
		restaurantRatingApp: restaurantRatingApp,
	
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

func (s *Server) AddAddress(ctx context.Context, r *pb.AddAddressRequest)(*pb.AddAddressResponse, error){
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

func (s *Server) RateRestaurant(ctx context.Context, r *pb.RateRestaurantRequest)(*pb.RateRestaurantResponse,error){
	return s.restaurantRatingApp.RateRestaurant(ctx,r)
}

func (s *Server) UpdateRestaurantRating(ctx context.Context, r *pb.UpdateRestaurantRatingRequest)(*pb.UpdateRestaurantRatingResponse,error){
	return s.restaurantRatingApp.UpdateRestaurantRating(ctx,r)
}

func (s *Server) ListRestaurantRatingByEater(ctx context.Context, r *pb.ListRestaurantRatingByEaterRequest)(*pb.ListRestaurantRatingByEaterResponse,error){
	return s.restaurantRatingApp.ListRestaurantRatingByEaterId(ctx,r)
}
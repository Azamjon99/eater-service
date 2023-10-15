package grpc

import (
	"context"
	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	service "github.com/Azamjon99/eater-service/src/application/services"
)
type Server struct{
	pb.EaterServiceServer
	eaterApp service.EaterApplicationService
	addressApp service.AddressApplicationService
	restaurantRatingApp service.RatingApplicationService
	walletApp service.WalletApplicationService
	orderApp service.OrderApplicationService

}


func NewServer(
	eaterApp service.EaterApplicationService,
	addressApp service.AddressApplicationService,
	restaurantRatingApp service.RatingApplicationService,
	walletApp service.WalletApplicationService,
	orderApp service.OrderApplicationService,

) *Server{
	return &Server{
		eaterApp: eaterApp,
		addressApp: addressApp,
		restaurantRatingApp: restaurantRatingApp,
		walletApp: walletApp,
		orderApp: orderApp,
	}
}

func (s *Server) SignupEater(ctx context.Context, r *pb.SignupEaterRequest)(*pb.SignupEaterResponse,error){
	return s.eaterApp.SignupEater(ctx,r)
}

func (s *Server) ConfirmSmsCode(ctx context.Context, r *pb.ConfirmSmsCodeRequest)(*pb.ConfirmSmsCodeResponse,error){
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

func (s *Server) DeleteAddress(ctx context.Context, r *pb.DeleteAddressRequest)(*pb.DeleteAddressResponse, error){
	return s.addressApp.DeleteAddress(ctx,r)
}

func (s *Server) ListAddressByEater(ctx context.Context, r *pb.ListAddressByEaterRequest)(*pb.ListAddressByEaterResponse, error){
	return s.addressApp.ListAddressByEaterId(ctx,r)
}

func (s *Server) GetAddress(ctx context.Context, r *pb.GetAddressRequest)(*pb.GetAddressResponse, error){
	return s.addressApp.GetAddressById(ctx,r)
}

func (s *Server) RateRestaurant(ctx context.Context, r *pb.RateRestaurantRequest)(*pb.RateRestaurantResponse,error){
	return s.restaurantRatingApp.RateRestaurant(ctx,r)
}

func (s *Server) UpdateRestaurantRating(ctx context.Context, r *pb.UpdateRestaurantRatingRequest)(*pb.UpdateRestaurantRatingResponse,error){
	return s.restaurantRatingApp.UpdateRestaurantRating(ctx,r)
}

func (s *Server) ListRestaurantRatingByEater(ctx context.Context, r *pb.ListRestaurantRatingByEaterRequest)(*pb.ListRestaurantRatingByEaterResponse,error){
	return s.restaurantRatingApp.ListDeliveryRatingByEaterId(ctx,r)
}
func (s *Server) RateDelivery(ctx context.Context, r *pb.RateDeliveryRequest)(*pb.RateDeliveryResponse,error){
	return s.restaurantRatingApp.RateDelivery(ctx,r)
}

func (s *Server) UpdateDeliveryRating(ctx context.Context, r *pb.UpdateDeliveryRatingRequest)(*pb.UpdateDeliveryRatingResponse,error){
	return s.restaurantRatingApp.UpdateDeliveryRating(ctx,r)
}

func (s *Server) AddPaymentCard(ctx context.Context, r *pb.AddPaymentCardRequest)(*pb.AddPaymentCardResponse,error){
	return s.walletApp.AddCard(ctx,r)
}

func (s *Server) GetPaymentCard(ctx context.Context, r *pb.GetPaymentCardRequest)(*pb.GetPaymentCardResponse,error){
	return s.walletApp.GetCard(ctx,r)
}
func (s *Server) ListPaymentCardByEater(ctx context.Context, r *pb.ListPaymentCardByEaterRequest)(*pb.ListPaymentCardByEaterResponse,error){
	return s.walletApp.ListCardByEaterId(ctx,r)
}

func (s *Server) DeletePaymentCard(ctx context.Context, r *pb.DeletePaymentCardRequest)(*pb.DeletePaymentCardResponse,error){
	return s.walletApp.DeleteCard(ctx,r)

}
func (s *Server) DeleteOrder(ctx context.Context, r *pb.DeleteOrderRequest)(*pb.DeleteOrderResponse,error){
	return s.orderApp.DeleteOrder(ctx,r)

}
func (s *Server) GetOrder(ctx context.Context, r *pb.GetOrderRequest)(*pb.GetOrderResponse,error){
	return s.orderApp.GetOrderById(ctx,r)

}

func (s *Server) ListOrderByEater(ctx context.Context, r *pb.ListOrderByEaterRequest)(*pb.ListOrderByEaterResponse,error){
	return s.orderApp.ListOrderByEaterId(ctx,r)

}
func (s *Server) PlaceOrder(ctx context.Context, r *pb.PlaceOrderRequest)(*pb.PlaceOrderResponse,error){
	return s.orderApp.PlaceOrder(ctx,r)

}

func (s *Server) UpdateOrder(ctx context.Context, r *pb.UpdateOrderRequest)(*pb.UpdateOrderResponse,error){
	return s.orderApp.UpdateOrder(ctx,r)

}
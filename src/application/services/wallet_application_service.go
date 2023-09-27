package service

import (
	"context"

	walletsvc "github.com/Azamjon99/eater-service/src/domain/wallet/services"
	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
)

type WalletApplicationService interface {
	AddCard(ctx context.Context, request *pb.AddPaymentCardRequest) (*pb.AddPaymentCardResponse, error)
	DeleteCard(ctx context.Context, request *pb.DeletePaymentCardRequest) (*pb.DeletePaymentCardResponse, error)
	GetCard(ctx context.Context, request *pb.GetPaymentCardRequest) (*pb.GetPaymentCardResponse, error)
	ListCardByEaterId(ctx context.Context, request *pb.ListPaymentCardByEaterRequest) (*pb.ListPaymentCardByEaterResponse, error)
}

type walletAppSvcImpl struct {
	walletSvc walletsvc.WalletService
}

func NewWalletApplicationService(walletSvc walletsvc.WalletService) WalletApplicationService {
	return &walletAppSvcImpl{
		walletSvc: walletSvc,
	}
}

func (s *walletAppSvcImpl) AddCard(ctx context.Context, request *pb.AddPaymentCardRequest) (*pb.AddPaymentCardResponse, error) {

	paymentCard, err := s.walletSvc.AddCard(ctx, request.EaterId, request.CardNumber, request.CardToken)
	if err != nil {
		return nil, err
	}

	return &pb.AddPaymentCardResponse{
		Card: &pb.PaymentCardView{
			Id:          paymentCard.ID,
			EaterId:     paymentCard.EaterID,
			Number:      paymentCard.Number,
		},
	}, nil
}

func (s *walletAppSvcImpl) DeleteCard(ctx context.Context, request *pb.DeletePaymentCardRequest) (*pb.DeletePaymentCardResponse, error) {

	_, err := s.walletSvc.DeleteCard(ctx, request.CardId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePaymentCardResponse{}, nil
}

func (s *walletAppSvcImpl) GetCard(ctx context.Context, request *pb.GetPaymentCardRequest) (*pb.GetPaymentCardResponse, error) {

	paymentCard, err := s.walletSvc.GetCard(ctx, request.CardId)
	if err != nil {
		return nil, err
	}

	return &pb.GetPaymentCardResponse{
		Card: &pb.PaymentCardView{
			Id:          paymentCard.ID,
			EaterId:     paymentCard.EaterID,
			Number:      paymentCard.Number,
		},
	}, nil
}

func (s *walletAppSvcImpl) ListCardByEaterId(ctx context.Context, request *pb.ListPaymentCardByEaterRequest) (*pb.ListPaymentCardByEaterResponse, error) {

	paymentCards, err := s.walletSvc.ListCardByEaterId(ctx, request.EaterId, request.Sort, int(request.Page), int(request.PageSize))
	if err != nil {
		return nil, err
	}

	var cardViews []*pb.PaymentCardView
	for _, paymentCard := range paymentCards {
		cardViews = append(cardViews, &pb.PaymentCardView{
			Id:          paymentCard.ID,
			EaterId:     paymentCard.EaterID,
			Number:      paymentCard.Number,
		})
	}

	return &pb.ListPaymentCardByEaterResponse{
		Cards: cardViews,
	}, nil
}
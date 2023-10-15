package service

import (
	"context"

	cardadded "github.com/Azamjon99/eater-service/src/application/integration_events/events/card_added"
	carddeleted "github.com/Azamjon99/eater-service/src/application/integration_events/events/card_deleted"
	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	walletsvc "github.com/Azamjon99/eater-service/src/domain/wallet/services"
	"github.com/Azamjon99/eater-service/src/infrastructure/pubsub"
	"go.uber.org/zap"
)

type WalletApplicationService interface {
	AddCard(ctx context.Context, request *pb.AddPaymentCardRequest) (*pb.AddPaymentCardResponse, error)
	DeleteCard(ctx context.Context, request *pb.DeletePaymentCardRequest) (*pb.DeletePaymentCardResponse, error)
	GetCard(ctx context.Context, request *pb.GetPaymentCardRequest) (*pb.GetPaymentCardResponse, error)
	ListCardByEaterId(ctx context.Context, request *pb.ListPaymentCardByEaterRequest) (*pb.ListPaymentCardByEaterResponse, error)
}

type walletAppSvcImpl struct {
	walletSvc walletsvc.WalletService
	producer pubsub.Producer
    logger *zap.Logger
}

func NewWalletApplicationService(walletSvc walletsvc.WalletService, producer pubsub.Producer,logger *zap.Logger) WalletApplicationService {
	return &walletAppSvcImpl{
		walletSvc: walletSvc,
		producer: producer,
        logger: logger,
	}
}

func (s *walletAppSvcImpl) AddCard(ctx context.Context, request *pb.AddPaymentCardRequest) (*pb.AddPaymentCardResponse, error) {

	paymentCard, err := s.walletSvc.AddCard(ctx, request.EaterId, request.CardNumber, request.CardToken)
	if err != nil {
		return nil, err
	}
	event := cardadded.NewEvent(paymentCard)
	if err := s.producer.PublishWithKey(event.Topic(),paymentCard.ID,event.Payload(),nil);err!=nil{
		s.logger.Error("error pushing message",zap.Error(err),zap.String("topic",event.Topic()))
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
	event := carddeleted.NewEvent(request.CardId)
	if err := s.producer.PublishWithKey(event.Topic(),request.CardId,event.Payload(),nil);err!=nil{
		s.logger.Error("error pushing message",zap.Error(err),zap.String("topic",event.Topic()))
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
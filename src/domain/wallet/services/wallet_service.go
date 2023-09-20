package services

import (
	"context"
	"eater-service/src/domain/wallet/models"
	"eater-service/src/domain/wallet/repositories"
	"eater-service/src/infrastructure/rand"
	"time"
)

type WalletService interface {
	AddCard(ctx context.Context, eaterId, number, cardToken, restaurant string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardId string) (*models.PaymentCard, error)
	GetCard(ctx context.Context, cardId string) (*models.PaymentCard, error)
	ListCardByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.PaymentCard, error)
}
type walletSvcImpl struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(walletRepo repositories.WalletRepository) WalletService {
	return &walletSvcImpl{
		walletRepo: walletRepo,
	}
}

func (s *walletSvcImpl) AddCard(ctx context.Context, eaterId, number, cardToken, restaurant string) (*models.PaymentCard, error) {

	paymentCard := &models.PaymentCard{
		ID:         rand.UUID(),
		EaterID:    eaterId,
		Number:     number,
		CardToken:  cardToken,
		Restaurant: restaurant,
		CreatedAt:  time.Now().UTC(),
	}

	err := s.walletRepo.AddCard(ctx, paymentCard); 
	if err != nil {
		return nil, err
	}

	return paymentCard, nil

}

func (s *walletSvcImpl) DeleteCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {

	err := s.walletRepo.DeleteCard(ctx, cardID); 
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *walletSvcImpl) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {

	card, err := s.walletRepo.GetCardInfo(ctx, cardID)
	if err != nil {
		return nil, err
	}

	return card, nil

}

func (s *walletSvcImpl) ListCardByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.PaymentCard, error) {

	card, err := s.walletRepo.ListCardByEaterId(ctx, eaterID, sort, page, pageSize)
	if err != nil {
		return nil, err
	}

	return card, nil
}

package repositories

import (
	"context"
	"eater-service/src/domain/wallet/models"
)

type WalletRepository interface {
	WithTx(ctx context.Context, f func(r WalletRepository) error) error
	AddCard(ctx context.Context, payment_card *models.PaymentCard) error
	DeleteCard(ctx context.Context, cardId string) error
	GetCardInfo(ctx context.Context, cardId string) (*models.PaymentCard, error)
	ListCardByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.PaymentCard, error)
}

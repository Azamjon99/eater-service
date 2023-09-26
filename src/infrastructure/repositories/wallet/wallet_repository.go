package wallet

import (
	"context"

	"github.com/Azamjon99/eater-service/src/domain/wallet/models"
	"github.com/Azamjon99/eater-service/src/domain/wallet/repositories"
	"github.com/Azamjon99/eater-service/src/infrastructure/repositories/utils"
	"gorm.io/gorm"
)

const (
	tablePaymentCard = "wallet.payment_card"
)

type walletRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.WalletRepository {
	return &walletRepoImpl{
		db: db,
	}
}

func (r *walletRepoImpl) AddCard(ctx context.Context, wallet *models.PaymentCard) error {
	result := r.db.WithContext(ctx).Table(tablePaymentCard).Create(wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *walletRepoImpl) DeleteCard(ctx context.Context, cardId string) error {
	var wallet models.PaymentCard
	result := r.db.WithContext(ctx).Table(tablePaymentCard).Delete(&wallet, "id = ?", cardId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *walletRepoImpl) GetCardInfo(ctx context.Context, cardId string) (*models.PaymentCard, error) {
	var wallet models.PaymentCard
	result := r.db.WithContext(ctx).Table(tablePaymentCard).First(&wallet, "id = ?", cardId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallet, nil
}

func (r *walletRepoImpl) ListCardByEaterId(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.PaymentCard, error) {
	var wallets []*models.PaymentCard
	result := r.db.WithContext(ctx).Table(tablePaymentCard).Where("eater_id = ?", eaterID)
	result.Scopes(utils.SortByCreatedAt(page, pageSize), utils.Sort(sort)).Find(&wallets)

	if result.Error != nil {
		return nil, result.Error
	}
	return wallets, nil
}

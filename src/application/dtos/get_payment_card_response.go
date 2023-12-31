package dtos

import "github.com/Azamjon99/eater-service/src/domain/wallet/models"

type GetPaymentCardResponse struct {
	PaymentCard *models.PaymentCard `json:"payment_card"`
}

func NewGetPaymentCardResponse(paymentCard *models.PaymentCard) *GetPaymentCardResponse {
	return &GetPaymentCardResponse{
		PaymentCard: paymentCard,
	}
}

package dtos

import "github.com/Azamjon99/eater-service/src/domain/wallet/models"

type ListPaymentCardResponse struct {
	PaymentCard []*models.PaymentCard `json:"payment_card"`
}

func NewListPaymentCardResponse(paymentCard []*models.PaymentCard) *ListPaymentCardResponse {
	return &ListPaymentCardResponse{
		PaymentCard: paymentCard,
	}
}

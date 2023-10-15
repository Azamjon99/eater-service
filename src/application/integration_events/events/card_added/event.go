package cardadded

import (
	"encoding/json"

	"github.com/Azamjon99/eater-service/src/domain/wallet/models"
)

const (
	topic = "card.added"
)

type Event struct{
	Card *models.PaymentCard `json:"card"`
}

func (e Event) Topic() string{
	return topic
}

func (e Event) Payload() []byte{
	jsonByte,err := json.Marshal(e)
	if err != nil{
		return nil
	}
	return jsonByte
}

func NewEvent(card *models.PaymentCard) Event{
	return Event{
		Card: card,
	}
}
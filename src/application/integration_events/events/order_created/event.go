package ordercreated

import (
	"encoding/json"

	"github.com/Azamjon99/eater-service/src/domain/order/models"
)



type Event struct {
	Order *models.Order 	`json:"order"`
}

func (Event) Topic() string {
	return "order.created"
}

func (e Event) Payload() []byte {
	jsonBytes, err:=json.Marshal(e)
	if err !=nil {
		return nil
	}
	return jsonBytes
}

func NewEvent(order *models.Order) Event {
	return Event {
		Order: order,
	}
}
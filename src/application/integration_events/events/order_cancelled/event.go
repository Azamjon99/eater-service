package ordercanceled

import "encoding/json"

const (
	topic = "order.canceled"
)

type Event struct{
	OrderID string `json:"order_id"`
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

func NewEvent(orderID string) Event{
	return Event{
		OrderID: orderID,
	}
}
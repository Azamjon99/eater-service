package carddeleted

import (
	"encoding/json"
)

	const (
		topic = "card.deleted"
	)


	type Event struct{
		CardID string `json:"card_id"`
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

	func NewEvent(cardID string) Event{
		return Event{
			CardID: cardID,
		}
	}
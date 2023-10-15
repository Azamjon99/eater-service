package orderrated

import (
	"encoding/json"
)

const (
	topic = "order.rated"
)
type Eater struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
 }

type Event struct{
	 RatingID  string `json:"rating_id"`
	 OrderId  string `json:"order_id"`
	 Rating    int    `json:"rating"`
	 Comment   string `json:"comment"`
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

func NewEvent(RatingID,OrderId,Comment string,Rating int) Event{
	return Event{
		RatingID:RatingID,
		OrderId:OrderId,
		Rating:Rating,
		Comment:Comment,
	}
}
package restaurantratingupdated

import "encoding/json"

const (
	topic = "restaurant.rating.updated"
)

type Event struct{
	RatingID string `json:"rating_id"`
	Rating 	 int `json:"rating"`
	Comment  string `json:"comment"`
}

func(e Event) Topic()string{
	return topic
}

func (e Event) Payload()[]byte{
	jsonByte,err := json.Marshal(e)

	if err != nil {
		return nil
	}

	return jsonByte
}

func NewEvent(RatingID,Comment string, Rating int) Event{
	return Event{
		RatingID: RatingID,
		Rating: Rating,
		Comment: Comment,
	}
}

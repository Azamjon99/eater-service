package restaurantrated

import (
	"encoding/json"
)

const (
	topic = "restaurant.rated"
)
type Eater struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
 }

type Event struct{
	 RestaurantID  string `json:"restaurant_id"`
	 RatingID  string `json:"rating_id"`
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

func NewEvent(RestaurantID,RatingID,Comment string,Rating int) Event{
	return Event{
		RestaurantID:RestaurantID,
		RatingID:RatingID,
		Rating:Rating,
		Comment:Comment,
	}
}
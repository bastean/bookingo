package event

import (
	"encoding/json"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
)

var CreatedSucceededTypeRoutingKey = messages.NewRoutingKey(&messages.MessageRoutingKey{
	Module:    "booking",
	Version:   "1",
	Type:      messages.Type.Event,
	Aggregate: "booking",
	Event:     "created",
	Status:    messages.Status.Succeeded,
})

type CreatedSucceededAttributes struct {
	HotelId   string
	Id        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CheckIn   string
	CheckOut  string
	Room      string
	Currency  string
	Total     float32
}

type CreatedSucceeded struct {
	Attributes *CreatedSucceededAttributes
}

func NewCreatedSucceeded(event *CreatedSucceeded) (*messages.Message, error) {
	attributes, err := json.Marshal(event.Attributes)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewCreatedSucceeded",
			What:  "failure to create event message attributes",
			Why: errors.Meta{
				"Routing Key": CreatedSucceededTypeRoutingKey,
			},
			Who: err,
		})
	}

	return messages.NewMessage(CreatedSucceededTypeRoutingKey, attributes, messages.Meta{}), nil
}

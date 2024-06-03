package message

import (
	"encoding/json"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
)

var CreatedSucceededEventTypeRoutingKey = messages.NewRoutingKey(&messages.MessageRoutingKey{
	Module:    "booking",
	Version:   "1",
	Type:      messages.Type.Event,
	Aggregate: "booking",
	Event:     "created",
	Status:    messages.Status.Succeeded,
})

type CreatedSucceededEventAttributes struct {
	HotelID   string
	ID        string
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

func NewCreatedSucceededEvent(attributes *CreatedSucceededEventAttributes) (*messages.Message, error) {
	attributesJson, err := json.Marshal(attributes)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewCreatedSucceededEvent",
			What:  "failure to create an event message",
			Why: errors.Meta{
				"Routing Key": CreatedSucceededEventTypeRoutingKey,
			},
			Who: err,
		})
	}

	return messages.NewMessage(CreatedSucceededEventTypeRoutingKey, attributesJson, messages.Meta{}), nil
}

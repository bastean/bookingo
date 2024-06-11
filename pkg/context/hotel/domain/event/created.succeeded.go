package event

import (
	"encoding/json"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
)

var CreatedSucceededTypeRoutingKey = messages.NewRoutingKey(&messages.MessageRoutingKey{
	Module:    "hotel",
	Version:   "1",
	Type:      messages.Type.Event,
	Aggregate: "hotel",
	Event:     "created",
	Status:    messages.Status.Succeeded,
})

type CreatedSucceededAttributes struct {
	ID    string
	Name  string
	Email string
	Phone string
}

type CreatedSucceeded struct {
	Attributes *CreatedSucceededAttributes
}

func NewCreatedSucceeded(event *CreatedSucceeded) (*messages.Message, error) {
	attributes, err := json.Marshal(event.Attributes)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewCreatedSucceededEvent",
			What:  "failure to create event message attributes",
			Why: errors.Meta{
				"Routing Key": CreatedSucceededTypeRoutingKey,
			},
			Who: err,
		})
	}

	return messages.NewMessage(CreatedSucceededTypeRoutingKey, attributes, messages.Meta{}), nil
}

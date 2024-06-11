package send

import (
	"encoding/json"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/queues"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type CreatedSucceededEventConsumer struct {
	models.UseCase[*event.CreatedSucceeded, types.Empty]
	Queues []*queues.Queue
}

func (consumer *CreatedSucceededEventConsumer) SubscribedTo() []*queues.Queue {
	return consumer.Queues
}

func (consumer *CreatedSucceededEventConsumer) On(message *messages.Message) error {
	hotel := new(event.CreatedSucceeded)

	hotel.Attributes = new(event.CreatedSucceededAttributes)

	err := json.Unmarshal(message.Attributes, hotel.Attributes)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "On",
			What:  "failure to obtain message attributes",
			Why: errors.Meta{
				"Id":          message.Id,
				"Routing Key": message.Type,
				"Occurred On": message.OccurredOn,
			},
			Who: err,
		})
	}

	_, err = consumer.UseCase.Run(hotel)

	if err != nil {
		return errors.BubbleUp(err, "On")
	}

	return nil
}

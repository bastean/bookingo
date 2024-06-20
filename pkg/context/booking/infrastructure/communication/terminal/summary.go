package terminal

import (
	"fmt"

	"github.com/bastean/bookingo/pkg/context/booking/domain/event"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Summary struct {
	models.Logger
}

func (client *Summary) Submit(data any) error {
	booking, ok := data.(*event.CreatedSucceededAttributes)

	if !ok {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure in type assertion",
			Why: errors.Meta{
				"Expected": new(event.CreatedSucceededAttributes),
				"Actual":   data,
			},
		})
	}

	summary := fmt.Sprintf(
		`
		Hi %s %s, 
		
		Here is the summary of your booking:

		Check-in: %s
		Check-out: %s
		Room: %s

		Total(%s): %g
		`,
		booking.FirstName, booking.LastName,
		booking.CheckIn, booking.CheckOut,
		booking.Room,
		booking.Currency, booking.Total,
	)

	client.Logger.Info(summary)

	return nil
}

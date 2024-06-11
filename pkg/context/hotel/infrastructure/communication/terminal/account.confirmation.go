package terminal

import (
	"fmt"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type AccountConfirmation struct {
	models.Logger
	ServerURL string
}

func (client *AccountConfirmation) Submit(data any) error {
	hotel, ok := data.(*event.CreatedSucceededAttributes)

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

	confirmationLink := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", hotel.Name, client.ServerURL, hotel.ID)

	client.Logger.Info(confirmationLink)

	return nil
}

package send

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/event"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Send struct {
	models.Transport
}

func (send *Send) Run(booking *event.CreatedSucceeded) (types.Empty, error) {
	err := send.Transport.Submit(booking.Attributes)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

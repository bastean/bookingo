package send

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Send struct {
	models.Transport
}

func (send *Send) Run(hotel *event.CreatedSucceeded) (types.Empty, error) {
	err := send.Transport.Submit(hotel.Attributes)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

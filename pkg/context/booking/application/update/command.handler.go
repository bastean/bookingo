package update

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type CommandHandler struct {
	models.UseCase[*aggregate.Booking, types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	primitive := aggregate.BookingPrimitive(*command)

	booking, err := aggregate.NewBooking(&primitive)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(booking)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}

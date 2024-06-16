package update

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Update struct {
	model.Repository
}

func (update *Update) Run(booking *aggregate.Booking) (types.Empty, error) {
	bookingRegistered, err := update.Repository.Search(&model.RepositorySearchCriteria{
		Id: booking.Id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if booking.HotelId.Value() != bookingRegistered.HotelId.Value() {
		return nil, nil
	}

	err = update.Repository.Update(booking)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

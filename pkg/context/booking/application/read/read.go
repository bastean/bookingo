package read

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
)

type Read struct {
	model.Repository
}

func (read *Read) Run(input *Input) (*aggregate.Booking, error) {
	booking, err := read.Repository.Search(&model.RepositorySearchCriteria{
		ID: input.ID,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if input.HotelID.Value() != booking.HotelID.Value() {
		return nil, nil
	}

	return booking, nil
}

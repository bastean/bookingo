package delete

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Delete struct {
	model.Repository
}

func (delete *Delete) Run(input *Input) (*types.Empty, error) {
	booking, err := delete.Repository.Search(&model.RepositorySearchCriteria{
		ID: input.ID,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if input.HotelID.Value() != booking.HotelID.Value() {
		return nil, nil
	}

	err = delete.Repository.Delete(booking.ID)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

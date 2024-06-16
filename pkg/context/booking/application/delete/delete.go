package delete

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Delete struct {
	model.Repository
}

func (delete *Delete) Run(input *Input) (types.Empty, error) {
	booking, err := delete.Repository.Search(&model.RepositorySearchCriteria{
		Id: input.Id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if input.HotelId.Value() != booking.HotelId.Value() {
		return nil, nil
	}

	err = delete.Repository.Delete(booking.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

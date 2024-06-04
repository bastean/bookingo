package read

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Read struct {
	model.Repository
}

func (read *Read) Run(id models.ValueObject[string]) (*aggregate.Hotel, error) {
	hotel, err := read.Repository.Search(model.RepositorySearchCriteria{
		ID: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return hotel, nil
}

package delete

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Delete struct {
	model.Repository
	models.Hashing
}

func (delete *Delete) Run(input *Input) (types.Empty, error) {
	hotel, err := delete.Repository.Search(model.RepositorySearchCriteria{
		ID: input.ID,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = services.IsPasswordInvalid(delete.Hashing, hotel.Password.Value(), input.Password.Value())

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = delete.Repository.Delete(hotel.ID)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

package verify

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Verify struct {
	model.Repository
}

func (verify *Verify) Run(id models.ValueObject[string]) (*types.Empty, error) {
	hotelRegistered, err := verify.Repository.Search(model.RepositorySearchCriteria{
		ID: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if hotelRegistered.Verified.Value() {
		return nil, nil
	}

	hotelRegistered.Verified, err = valueobj.NewVerified(true)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	hotelRegistered.Password = nil

	err = verify.Repository.Update(hotelRegistered)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

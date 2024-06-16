package update

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Update struct {
	model.Repository
	models.Hashing
}

func (update *Update) Run(input *Input) (types.Empty, error) {
	hotelRegistered, err := update.Repository.Search(model.RepositorySearchCriteria{
		Id: input.Hotel.Id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = services.IsPasswordInvalid(update.Hashing, hotelRegistered.Password.Value(), input.Hotel.Password.Value())

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if input.UpdatedPassword != nil {
		input.Hotel.Password = input.UpdatedPassword
	}

	input.Hotel.Verified = hotelRegistered.Verified

	err = update.Repository.Update(input.Hotel)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

package login

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

type Login struct {
	model.Repository
	models.Hashing
}

func (login *Login) Run(input *Input) (*aggregate.Hotel, error) {
	hotel, err := login.Repository.Search(model.RepositorySearchCriteria{
		Email: input.Email,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = services.IsPasswordInvalid(login.Hashing, hotel.Password.Value(), input.Password.Value())

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return hotel, nil
}

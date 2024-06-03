package update

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Update struct {
	model.Repository
	models.Hashing
}

func (update *Update) Run(hotelUpdate *Command) (*types.Empty, error) {
	idVO, err := valueobj.NewId(hotelUpdate.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	hotelRegistered, err := update.Repository.Search(model.RepositorySearchCriteria{
		Id: idVO,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = services.IsPasswordInvalid(update.Hashing, hotelRegistered.Password.Value(), hotelUpdate.Password)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	var errEmail, errHotelname, errPassword error

	if hotelUpdate.Email != "" {
		hotelRegistered.Email, errEmail = valueobj.NewEmail(hotelUpdate.Email)
	}

	if hotelUpdate.Hotelname != "" {
		hotelRegistered.Hotelname, errHotelname = valueobj.NewHotelname(hotelUpdate.Hotelname)
	}

	if hotelUpdate.UpdatedPassword != "" {
		hotelRegistered.Password, errPassword = valueobj.NewPassword(hotelUpdate.UpdatedPassword)
	} else {
		hotelRegistered.Password = nil
	}

	err = errors.Join(errEmail, errHotelname, errPassword)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = update.Repository.Update(hotelRegistered)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

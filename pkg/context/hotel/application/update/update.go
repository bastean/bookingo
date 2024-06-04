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
	idVO, err := valueobj.NewId(hotelUpdate.ID)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	hotelRegistered, err := update.Repository.Search(model.RepositorySearchCriteria{
		ID: idVO,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = services.IsPasswordInvalid(update.Hashing, hotelRegistered.Password.Value(), hotelUpdate.Password)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	var errName, errEmail, errPhone, errPassword error

	if hotelUpdate.Name != "" {
		hotelRegistered.Name, errName = valueobj.NewName(hotelUpdate.Name)
	}

	if hotelUpdate.Email != "" {
		hotelRegistered.Email, errEmail = valueobj.NewEmail(hotelUpdate.Email)
	}

	if hotelUpdate.Phone != "" {
		hotelRegistered.Phone, errPhone = valueobj.NewPhone(hotelUpdate.Phone)
	}

	if hotelUpdate.UpdatedPassword != "" {
		hotelRegistered.Password, errPassword = valueobj.NewPassword(hotelUpdate.UpdatedPassword)
	} else {
		hotelRegistered.Password = nil
	}

	err = errors.Join(errName, errEmail, errPhone, errPassword)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = update.Repository.Update(hotelRegistered)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

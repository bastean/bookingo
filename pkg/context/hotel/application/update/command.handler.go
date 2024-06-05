package update

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Input struct {
	Hotel           *aggregate.Hotel
	UpdatedPassword models.ValueObject[string]
}

type CommandHandler struct {
	models.UseCase[*Input, *types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	hotel, err := aggregate.NewHotel(command.ID, command.Name, command.Email, command.Phone, command.Password)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	var updatedPassword models.ValueObject[string]

	if command.UpdatedPassword != "" {
		updatedPassword, err = valueobj.NewPassword(command.UpdatedPassword)

		if err != nil {
			return errors.BubbleUp(err, "Handle")
		}
	}

	_, err = handler.UseCase.Run(&Input{
		Hotel:           hotel,
		UpdatedPassword: updatedPassword,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}

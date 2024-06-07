package delete

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Input struct {
	HotelID, ID models.ValueObject[string]
}

type CommandHandler struct {
	models.UseCase[*Input, *types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	hotelID, errHotelID := valueobj.NewId(command.HotelID)
	id, errID := valueobj.NewId(command.ID)

	err := errors.Join(errHotelID, errID)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(&Input{
		HotelID: hotelID,
		ID:      id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}

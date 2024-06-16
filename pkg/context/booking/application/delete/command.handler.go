package delete

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Input struct {
	HotelId, Id models.ValueObject[string]
}

type CommandHandler struct {
	models.UseCase[*Input, types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	hotelId, errHotelId := valueobj.NewId(command.HotelId)
	id, errId := valueobj.NewId(command.Id)

	err := errors.Join(errHotelId, errId)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(&Input{
		HotelId: hotelId,
		Id:      id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}

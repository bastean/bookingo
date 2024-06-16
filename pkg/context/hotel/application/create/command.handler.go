package create

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type CommandHandler struct {
	models.UseCase[*aggregate.Hotel, types.Empty]
	models.Broker
}

func (handler *CommandHandler) Handle(command *Command) error {
	hotel, err := aggregate.NewHotel(&aggregate.HotelPrimitive{
		Id:       command.Id,
		Name:     command.Name,
		Email:    command.Email,
		Phone:    command.Phone,
		Password: command.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(hotel)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	handler.Broker.PublishMessages(hotel.PullMessages())

	return nil
}

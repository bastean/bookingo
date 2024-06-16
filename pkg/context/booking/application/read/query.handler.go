package read

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Input struct {
	HotelId, Id models.ValueObject[string]
}

type QueryHandler struct {
	models.UseCase[*Input, *aggregate.Booking]
}

func (handler *QueryHandler) Handle(query *Query) (*Response, error) {
	hotelId, errHotelId := valueobj.NewId(query.HotelId)
	id, errId := valueobj.NewId(query.Id)

	err := errors.Join(errHotelId, errId)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	booking, err := handler.UseCase.Run(&Input{
		HotelId: hotelId,
		Id:      id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*booking.ToPrimitives())

	return &response, nil
}

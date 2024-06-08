package read

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Input struct {
	HotelID, ID models.ValueObject[string]
}

type QueryHandler struct {
	models.UseCase[*Input, *aggregate.Booking]
}

func (handler *QueryHandler) Handle(query *Query) (*Response, error) {
	hotelID, errHotelID := valueobj.NewId(query.HotelID)
	id, errID := valueobj.NewId(query.ID)

	err := errors.Join(errHotelID, errID)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	booking, err := handler.UseCase.Run(&Input{
		HotelID: hotelID,
		ID:      id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*booking.ToPrimitives())

	return &response, nil
}

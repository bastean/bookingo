package read

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type QueryHandler struct {
	models.UseCase[models.ValueObject[string], *aggregate.Hotel]
}

func (handler *QueryHandler) Handle(query *Query) (*Response, error) {
	id, err := valueobj.NewId(query.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	hotel, err := handler.UseCase.Run(id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*hotel.ToPrimitives())

	return &response, nil
}

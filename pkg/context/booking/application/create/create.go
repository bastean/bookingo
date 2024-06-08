package create

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Create struct {
	model.Repository
}

func (create *Create) Run(booking *aggregate.Booking) (types.Empty, error) {
	err := create.Repository.Save(booking)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

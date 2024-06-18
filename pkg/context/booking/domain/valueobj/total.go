package valueobj

import (
	"fmt"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Total struct {
	Total float32
}

func (total *Total) Value() float32 {
	return total.Total
}

func (total *Total) IsValid() error {
	return nil
}

func NewTotal(total float32) (models.ValueObject[float32], error) {
	totalVO := &Total{
		Total: total,
	}

	if totalVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewTotal",
			What:  "invalid total value",
			Why: errors.Meta{
				"Total": fmt.Sprintf("%s", total),
			},
		})
	}

	return totalVO, nil
}

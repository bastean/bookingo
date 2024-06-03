package valueobj

import (
	"strings"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

type Check struct {
	Check string `validate:"datetime=02-01-2006"`
}

func (check *Check) Value() string {
	return check.Check
}

func (check *Check) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(check)
}

func NewCheck(check string) (models.ValueObject[string], error) {
	check = strings.TrimSpace(check)

	checkVO := &Check{
		Check: check,
	}

	if checkVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewCheck",
			What:  "invalid check format, required e.g. DD-MM-YYYY",
			Why: errors.Meta{
				"Check": check,
			},
		})
	}

	return checkVO, nil
}

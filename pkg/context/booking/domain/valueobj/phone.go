package valueobj

import (
	"strings"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

type Phone struct {
	Phone string `validate:"e164"`
}

func (phone *Phone) Value() string {
	return phone.Phone
}

func (phone *Phone) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(phone)
}

func NewPhone(phone string) (models.ValueObject[string], error) {
	phone = strings.TrimSpace(phone)

	phoneVO := &Phone{
		Phone: phone,
	}

	if phoneVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewPhone",
			What:  "invalid phone format, required e.g. +1123456789",
			Why: errors.Meta{
				"Phone": phone,
			},
		})
	}

	return phoneVO, nil
}

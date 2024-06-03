package valueobj

import (
	"strings"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

type Currency struct {
	Currency string `validate:"iso4217"`
}

func (currency *Currency) Value() string {
	return currency.Currency
}

func (currency *Currency) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(currency)
}

func NewCurrency(currency string) (models.ValueObject[string], error) {
	currency = strings.TrimSpace(currency)

	currencyVO := &Currency{
		Currency: currency,
	}

	if currencyVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewCurrency",
			What:  "invalid currency format, required e.g. USD",
			Why: errors.Meta{
				"Currency": currency,
			},
		})
	}

	return currencyVO, nil
}

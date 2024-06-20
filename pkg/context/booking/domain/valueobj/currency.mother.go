package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

func RandomCurrency() models.ValueObject[string] {
	// TODO: services.Create.CurrencyShort()
	value, err := NewCurrency("USD")

	if err != nil {
		errors.Panic(err.Error(), "RandomCurrency")
	}

	return value
}

func InvalidCurrency() (string, error) {
	value := "x"

	_, err := NewCurrency(value)

	return value, err
}

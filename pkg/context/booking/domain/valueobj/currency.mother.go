package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomCurrency() models.ValueObject[string] {
	value, err := NewCurrency(services.Create.CurrencyShort())

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

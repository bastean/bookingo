package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomCurrency() (models.ValueObject[string], error) {
	return NewCurrency(services.Create.CurrencyShort())
}

func InvalidCurrency() (string, error) {
	value := "x"

	_, err := NewCurrency(value)

	return value, err
}

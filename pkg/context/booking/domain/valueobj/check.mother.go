package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomCheck() (models.ValueObject[string], error) {
	return NewCheck(services.Create.Date().Format("02-01-2006"))
}

func InvalidCheck() (string, error) {
	value := "x"

	_, err := NewCheck(value)

	return value, err
}

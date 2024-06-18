package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomCheck() models.ValueObject[string] {
	value, err := NewCheck(services.Create.Date().Format("02-01-2006"))

	if err != nil {
		errors.Panic(err.Error(), "RandomCheck")
	}

	return value
}

func InvalidCheck() (string, error) {
	value := "x"

	_, err := NewCheck(value)

	return value, err
}

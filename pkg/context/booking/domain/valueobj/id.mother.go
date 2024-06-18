package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomId() models.ValueObject[string] {
	value, err := NewId(services.Create.UUID())

	if err != nil {
		errors.Panic(err.Error(), "RandomId")
	}

	return value
}

func InvalidId() (string, error) {
	value := "x"

	_, err := NewId(value)

	return value, err
}

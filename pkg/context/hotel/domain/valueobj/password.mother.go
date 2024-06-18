package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomPassword() models.ValueObject[string] {
	value, err := NewPassword(services.Create.Regex(`[\W\w]{8,64}`))

	if err != nil {
		errors.Panic(err.Error(), "RandomName")
	}

	return value
}

func WithInvalidPasswordLength() (string, error) {
	value := "x"

	_, err := NewPassword(value)

	return value, err
}

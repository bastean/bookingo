package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomPassword() (models.ValueObject[string], error) {
	return NewPassword(services.Create.Regex(`[\W\w]{8,64}`))
}

func WithInvalidPasswordLength() (string, error) {
	value := "x"

	_, err := NewPassword(value)

	return value, err
}

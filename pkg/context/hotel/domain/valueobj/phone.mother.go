package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomPhone() (models.ValueObject[string], error) {
	return NewPhone(services.Create.Phone())
}

func InvalidPhone() (string, error) {
	value := "x"

	_, err := NewPhone(value)

	return value, err
}

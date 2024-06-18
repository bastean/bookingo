package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomPhone() models.ValueObject[string] {
	value, err := NewPhone("+" + services.Create.Phone())

	if err != nil {
		errors.Panic(err.Error(), "RandomPhone")
	}

	return value
}

func InvalidPhone() (string, error) {
	value := "x"

	_, err := NewPhone(value)

	return value, err
}

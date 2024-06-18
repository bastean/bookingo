package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomEmail() models.ValueObject[string] {
	value, err := NewEmail(services.Create.Email())

	if err != nil {
		errors.Panic(err.Error(), "RandomEmail")
	}

	return value
}

func InvalidEmail() (string, error) {
	value := "x"

	_, err := NewEmail(value)

	return value, err
}

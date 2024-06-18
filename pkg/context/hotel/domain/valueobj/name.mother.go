package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomName() models.ValueObject[string] {
	value, err := NewName(services.Create.Regex(`^[\w\s-]{1,64}$`))

	if err != nil {
		errors.Panic(err.Error(), "RandomName")
	}

	return value
}

func WithInvalidNameLength() (string, error) {
	value := ""

	_, err := NewName(value)

	return value, err
}

func WithInvalidNameAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewName(value)

	return value, err
}

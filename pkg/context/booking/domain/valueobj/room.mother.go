package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomRoom() models.ValueObject[string] {
	value, err := NewRoom(services.Create.Regex(`^[\w\s-]{1,64}$`))

	if err != nil {
		errors.Panic(err.Error(), "RandomRoom")
	}

	return value
}

func WithInvalidRoomLength() (string, error) {
	value := ""

	_, err := NewRoom(value)

	return value, err
}

func WithInvalidRoomAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewRoom(value)

	return value, err
}

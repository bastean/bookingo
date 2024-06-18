package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomTotal() models.ValueObject[float32] {
	value, err := NewTotal(services.Create.Float32())

	if err != nil {
		errors.Panic(err.Error(), "RandomRoom")
	}

	return value
}

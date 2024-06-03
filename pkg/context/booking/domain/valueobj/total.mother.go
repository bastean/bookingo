package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/services"
)

func RandomTotal() (models.ValueObject[float32], error) {
	return NewTotal(services.Create.Float32())
}

package valueobj

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/valueobjs"
)

func NewId(id string) (models.ValueObject[string], error) {
	return valueobjs.NewId(id)
}

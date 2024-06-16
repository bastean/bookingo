package model

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type RepositorySearchCriteria struct {
	Id    models.ValueObject[string]
	Email models.ValueObject[string]
	Phone models.ValueObject[string]
}

type Repository interface {
	Save(hotel *aggregate.Hotel) error
	Update(hotel *aggregate.Hotel) error
	Verify(id models.ValueObject[string]) error
	Delete(id models.ValueObject[string]) error
	Search(criteria *RepositorySearchCriteria) (*aggregate.Hotel, error)
}

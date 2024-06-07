package model

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type RepositorySearchCriteria struct {
	HotelID models.ValueObject[string]
	ID      models.ValueObject[string]
	Email   models.ValueObject[string]
	Phone   models.ValueObject[string]
}

type Repository interface {
	Save(booking *aggregate.Booking) error
	Update(booking *aggregate.Booking) error
	Delete(id models.ValueObject[string]) error
	Search(criteria *RepositorySearchCriteria) (*aggregate.Booking, error)
}

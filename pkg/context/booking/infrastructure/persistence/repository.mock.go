package persistence

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (repository *RepositoryMock) Save(booking *aggregate.Booking) error {
	repository.Called(booking)
	return nil
}

func (repository *RepositoryMock) Update(booking *aggregate.Booking) error {
	repository.Called(booking)
	return nil
}

func (repository *RepositoryMock) Delete(id models.ValueObject[string]) error {
	repository.Called(id)
	return nil
}

func (repository *RepositoryMock) Search(criteria *model.RepositorySearchCriteria) (*aggregate.Booking, error) {
	args := repository.Called(criteria)
	return args.Get(0).(*aggregate.Booking), nil
}

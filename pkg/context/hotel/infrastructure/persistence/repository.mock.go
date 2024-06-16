package persistence

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (repository *RepositoryMock) Save(hotel *aggregate.Hotel) error {
	repository.Called(hotel)
	return nil
}

func (repository *RepositoryMock) Update(hotel *aggregate.Hotel) error {
	repository.Called(hotel)
	return nil
}

func (repository *RepositoryMock) Verify(id models.ValueObject[string]) error {
	repository.Called(id)
	return nil
}

func (repository *RepositoryMock) Delete(id models.ValueObject[string]) error {
	repository.Called(id)
	return nil
}

func (repository *RepositoryMock) Search(criteria *model.RepositorySearchCriteria) (*aggregate.Hotel, error) {
	args := repository.Called(criteria)
	return args.Get(0).(*aggregate.Hotel), nil
}

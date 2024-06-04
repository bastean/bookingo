package read_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/application/read"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/stretchr/testify/suite"
)

type HotelReadTestSuite struct {
	suite.Suite
	sut        models.QueryHandler[*read.Query, *read.Response]
	useCase    models.UseCase[models.ValueObject[string], *aggregate.Hotel]
	repository *persistence.RepositoryMock
}

func (suite *HotelReadTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &read.Read{
		Repository: suite.repository,
	}

	suite.sut = &read.QueryHandler{
		UseCase: suite.useCase,
	}
}

func (suite *HotelReadTestSuite) TestLogin() {
	hotel := aggregate.RandomHotel()

	query := &read.Query{
		ID: hotel.ID.Value(),
	}

	filter := model.RepositorySearchCriteria{
		ID: hotel.ID,
	}

	suite.repository.On("Search", filter).Return(hotel)

	expected := hotel.ToPrimitives()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitHotelReadSuite(t *testing.T) {
	suite.Run(t, new(HotelReadTestSuite))
}

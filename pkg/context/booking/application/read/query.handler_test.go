package read_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/application/read"
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/booking/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/stretchr/testify/suite"
)

type BookingReadTestSuite struct {
	suite.Suite
	sut        models.QueryHandler[*read.Query, *read.Response]
	useCase    models.UseCase[*read.Input, *aggregate.Booking]
	repository *persistence.RepositoryMock
}

func (suite *BookingReadTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &read.Read{
		Repository: suite.repository,
	}

	suite.sut = &read.QueryHandler{
		UseCase: suite.useCase,
	}
}

func (suite *BookingReadTestSuite) TestRead() {
	booking := aggregate.RandomBooking()

	query := &read.Query{
		HotelID: booking.HotelID.Value(),
		ID:      booking.ID.Value(),
	}

	criteria := &model.RepositorySearchCriteria{
		ID: booking.ID,
	}

	suite.repository.On("Search", criteria).Return(booking)

	expected := booking.ToPrimitives()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitBookingReadSuite(t *testing.T) {
	suite.Run(t, new(BookingReadTestSuite))
}

package delete_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/application/delete"
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/booking/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/stretchr/testify/suite"
)

type BookingDeleteTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*delete.Command]
	useCase    models.UseCase[*delete.Input, types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *BookingDeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &delete.Delete{
		Repository: suite.repository,
	}

	suite.sut = &delete.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *BookingDeleteTestSuite) TestDelete() {
	booking := aggregate.RandomBooking()

	command := &delete.Command{
		HotelId: booking.HotelId.Value(),
		Id:      booking.Id.Value(),
	}

	criteria := &model.RepositorySearchCriteria{
		Id: booking.Id,
	}

	suite.repository.On("Search", criteria).Return(booking)

	suite.repository.On("Delete", booking.Id)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitBookingDeleteSuite(t *testing.T) {
	suite.Run(t, new(BookingDeleteTestSuite))
}

package update_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/application/update"
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/booking/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/stretchr/testify/suite"
)

type BookingUpdateTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*update.Command]
	useCase    models.UseCase[*aggregate.Booking, *types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *BookingUpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &update.Update{
		Repository: suite.repository,
	}

	suite.sut = &update.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *BookingUpdateTestSuite) TestUpdate() {
	command := update.RandomCommand()

	primitive := aggregate.BookingPrimitive(*command)

	booking, _ := aggregate.NewBooking(&primitive)

	criteria := &model.RepositorySearchCriteria{
		ID: booking.ID,
	}

	suite.repository.On("Search", criteria).Return(booking)

	suite.repository.On("Update", booking)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitBookingUpdateSuite(t *testing.T) {
	suite.Run(t, new(BookingUpdateTestSuite))
}

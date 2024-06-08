package create_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/application/create"
	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/communications"
	"github.com/stretchr/testify/suite"
)

type BookingCreateTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*create.Command]
	useCase    models.UseCase[*aggregate.Booking, types.Empty]
	repository *persistence.RepositoryMock
	broker     *communications.BrokerMock
}

func (suite *BookingCreateTestSuite) SetupTest() {
	suite.broker = new(communications.BrokerMock)

	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &create.Create{
		Repository: suite.repository,
	}

	suite.sut = &create.CommandHandler{
		UseCase: suite.useCase,
		Broker:  suite.broker,
	}
}

func (suite *BookingCreateTestSuite) TestCreate() {
	command := create.RandomCommand()

	primitive := aggregate.BookingPrimitive(*command)

	booking, _ := aggregate.NewBooking(&primitive)

	messages := booking.Messages

	suite.repository.On("Save", booking)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitBookingCreateSuite(t *testing.T) {
	suite.Run(t, new(BookingCreateTestSuite))
}

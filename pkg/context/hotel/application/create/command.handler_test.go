package create_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/application/create"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/communications"
	"github.com/stretchr/testify/suite"
)

type HotelCreateTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*create.Command]
	useCase    models.UseCase[*aggregate.Hotel, types.Empty]
	repository *persistence.RepositoryMock
	broker     *communications.BrokerMock
}

func (suite *HotelCreateTestSuite) SetupTest() {
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

func (suite *HotelCreateTestSuite) TestCreate() {
	command := create.RandomCommand()

	hotel, _ := aggregate.NewHotel(command.Id, command.Name, command.Email, command.Phone, command.Password)

	messages := hotel.Messages

	suite.repository.On("Save", hotel)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitHotelCreateSuite(t *testing.T) {
	suite.Run(t, new(HotelCreateTestSuite))
}

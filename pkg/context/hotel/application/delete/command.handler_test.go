package delete_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/application/delete"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/cryptographics"
	"github.com/stretchr/testify/suite"
)

type HotelDeleteTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*delete.Command]
	useCase    models.UseCase[*delete.Input, *types.Empty]
	hashing    *cryptographics.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *HotelDeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographics.HashingMock)

	suite.useCase = &delete.Delete{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &delete.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *HotelDeleteTestSuite) TestDelete() {
	hotel := aggregate.RandomHotel()

	command := &delete.Command{
		Id:       hotel.Id.Value(),
		Password: hotel.Password.Value(),
	}

	filter := model.RepositorySearchCriteria{
		Id: hotel.Id,
	}

	suite.repository.On("Search", filter).Return(hotel)

	suite.hashing.On("IsNotEqual", hotel.Password.Value(), hotel.Password.Value()).Return(false)

	suite.repository.On("Delete", hotel.Id)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitHotelDeleteSuite(t *testing.T) {
	suite.Run(t, new(HotelDeleteTestSuite))
}

package update_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/application/update"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/cryptographics"
	"github.com/stretchr/testify/suite"
)

type HotelUpdateTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*update.Command]
	useCase    models.UseCase[*update.Input, types.Empty]
	hashing    *cryptographics.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *HotelUpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographics.HashingMock)

	suite.useCase = &update.Update{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &update.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *HotelUpdateTestSuite) TestUpdate() {
	command := update.RandomCommand()

	hotel, _ := aggregate.NewHotel(&aggregate.HotelPrimitive{
		Id:       command.Id,
		Name:     command.Name,
		Email:    command.Email,
		Phone:    command.Phone,
		Password: command.UpdatedPassword,
	})

	idVO, _ := valueobj.NewId(command.Id)

	criteria := &model.RepositorySearchCriteria{
		Id: idVO,
	}

	suite.repository.On("Search", criteria).Return(hotel)

	suite.hashing.On("IsNotEqual", hotel.Password.Value(), command.Password).Return(false)

	suite.repository.On("Update", hotel)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUnitHotelUpdateSuite(t *testing.T) {
	suite.Run(t, new(HotelUpdateTestSuite))
}

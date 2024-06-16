package verify_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/application/verify"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/stretchr/testify/suite"
)

type HotelVerifyTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*verify.Command]
	useCase    models.UseCase[models.ValueObject[string], types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *HotelVerifyTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &verify.Verify{
		Repository: suite.repository,
	}

	suite.sut = &verify.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *HotelVerifyTestSuite) TestVerify() {
	command := verify.RandomCommand()

	hotel := aggregate.RandomHotel()

	idVO, _ := valueobj.NewId(command.Id)

	hotel.Id = idVO

	criteria := &model.RepositorySearchCriteria{
		Id: idVO,
	}

	suite.repository.On("Search", criteria).Return(hotel)

	suite.repository.On("Verify", idVO)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitHotelVerifySuite(t *testing.T) {
	suite.Run(t, new(HotelVerifyTestSuite))
}

package login_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/application/login"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/cryptographics"
	"github.com/stretchr/testify/suite"
)

type HotelLoginTestSuite struct {
	suite.Suite
	sut        models.QueryHandler[*login.Query, *login.Response]
	useCase    models.UseCase[*login.Input, *aggregate.Hotel]
	hashing    *cryptographics.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *HotelLoginTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographics.HashingMock)

	suite.useCase = &login.Login{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &login.QueryHandler{
		UseCase: suite.useCase,
	}
}

func (suite *HotelLoginTestSuite) TestLogin() {
	hotel := aggregate.RandomHotel()

	query := &login.Query{
		Email:    hotel.Email.Value(),
		Password: hotel.Password.Value(),
	}

	filter := model.RepositorySearchCriteria{
		Email: hotel.Email,
	}

	suite.repository.On("Search", filter).Return(hotel)

	suite.hashing.On("IsNotEqual", hotel.Password.Value(), hotel.Password.Value()).Return(false)

	expected := hotel.ToPrimitives()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitHotelLoginSuite(t *testing.T) {
	suite.Run(t, new(HotelLoginTestSuite))
}

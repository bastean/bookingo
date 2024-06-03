package persistence_test

import (
	"os"
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/cryptographics"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/persistences"
	"github.com/stretchr/testify/suite"
)

type HotelMongoRepositoryTestSuite struct {
	suite.Suite
	sut     model.Repository
	hashing *cryptographics.HashingMock
}

func (suite *HotelMongoRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_URI")

	databaseName := "bookingo-test"

	database, _ := persistences.NewMongoDatabase(uri, databaseName)

	collectionName := "hotels-test"

	suite.hashing = new(cryptographics.HashingMock)

	suite.sut, _ = persistence.NewMongoCollection(database, collectionName, suite.hashing)
}

func (suite *HotelMongoRepositoryTestSuite) TestSave() {
	hotel := aggregate.RandomHotel()

	suite.hashing.On("Hash", hotel.Password.Value()).Return(hotel.Password.Value())

	suite.NoError(suite.sut.Save(hotel))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *HotelMongoRepositoryTestSuite) TestSaveDuplicate() {
	hotel := aggregate.RandomHotel()

	suite.hashing.On("Hash", hotel.Password.Value()).Return(hotel.Password.Value())

	suite.NoError(suite.sut.Save(hotel))

	suite.Error(suite.sut.Save(hotel))
}

func (suite *HotelMongoRepositoryTestSuite) TestUpdate() {
	hotel := aggregate.RandomHotel()

	suite.hashing.On("Hash", hotel.Password.Value()).Return(hotel.Password.Value())

	suite.NoError(suite.sut.Save(hotel))

	password, _ := valueobj.RandomPassword()

	hotel.Password = password

	suite.hashing.On("Hash", hotel.Password.Value()).Return(hotel.Password.Value())

	suite.NoError(suite.sut.Update(hotel))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *HotelMongoRepositoryTestSuite) TestDelete() {
	hotel := aggregate.RandomHotel()

	suite.hashing.On("Hash", hotel.Password.Value()).Return(hotel.Password.Value())

	suite.NoError(suite.sut.Save(hotel))

	suite.NoError(suite.sut.Delete(hotel.Id))
}

func (suite *HotelMongoRepositoryTestSuite) TestSearch() {
	expected := aggregate.RandomHotel()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value()).Return(expected.Password.Value())

	suite.NoError(suite.sut.Save(expected))

	filter := model.RepositorySearchCriteria{
		Email: expected.Email,
	}

	hotel, err := suite.sut.Search(filter)

	suite.NoError(err)

	actual := hotel

	suite.EqualValues(expected, actual)
}

func TestIntegrationHotelMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(HotelMongoRepositoryTestSuite))
}

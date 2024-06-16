package persistence_test

import (
	"os"
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/booking/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/persistences"
	"github.com/stretchr/testify/suite"
)

type BookingMongoRepositoryTestSuite struct {
	suite.Suite
	sut model.Repository
}

func (suite *BookingMongoRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_URI")

	databaseName := "bookingo-test"

	database, _ := persistences.NewMongoDatabase(uri, databaseName)

	collectionName := "bookings-test"

	suite.sut, _ = persistence.NewMongoCollection(database, collectionName)
}

func (suite *BookingMongoRepositoryTestSuite) TestSave() {
	booking := aggregate.RandomBooking()

	suite.NoError(suite.sut.Save(booking))
}

func (suite *BookingMongoRepositoryTestSuite) TestSaveDuplicate() {
	booking := aggregate.RandomBooking()

	suite.NoError(suite.sut.Save(booking))

	suite.Error(suite.sut.Save(booking))
}

func (suite *BookingMongoRepositoryTestSuite) TestUpdate() {
	booking := aggregate.RandomBooking()

	suite.NoError(suite.sut.Save(booking))

	bookingUpdated := aggregate.RandomBooking()

	bookingUpdated.HotelId = booking.HotelId

	bookingUpdated.Id = booking.Id

	suite.NoError(suite.sut.Update(booking))
}

func (suite *BookingMongoRepositoryTestSuite) TestDelete() {
	booking := aggregate.RandomBooking()

	suite.NoError(suite.sut.Save(booking))

	suite.NoError(suite.sut.Delete(booking.Id))
}

func (suite *BookingMongoRepositoryTestSuite) TestSearch() {
	expected := aggregate.RandomBooking()

	expected.PullMessages()

	suite.NoError(suite.sut.Save(expected))

	criteria := &model.RepositorySearchCriteria{
		Id: expected.Id,
	}

	booking, err := suite.sut.Search(criteria)

	suite.NoError(err)

	actual := booking

	suite.EqualValues(expected, actual)
}

func TestIntegrationBookingMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(BookingMongoRepositoryTestSuite))
}

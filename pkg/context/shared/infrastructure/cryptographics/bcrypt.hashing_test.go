package cryptographics_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/cryptographics"
	"github.com/stretchr/testify/suite"
)

type HotelBcryptHashingTestSuite struct {
	suite.Suite
	sut models.Hashing
}

func (suite *HotelBcryptHashingTestSuite) SetupTest() {
	suite.sut = new(cryptographics.Bcrypt)
}

func (suite *HotelBcryptHashingTestSuite) TestHash() {
	password, _ := valueobj.RandomPassword()

	plain := password.Value()

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	suite.NotEqual(plain, hashed)
}

func (suite *HotelBcryptHashingTestSuite) TestIsNotEqual() {
	password, _ := valueobj.RandomPassword()

	plain := password.Value()

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	isNotEqual := suite.sut.IsNotEqual(hashed, plain)

	suite.False(isNotEqual)
}

func TestIntegrationHotelBcryptHashingSuite(t *testing.T) {
	suite.Run(t, new(HotelBcryptHashingTestSuite))
}

package valueobj_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type RoomValueObjectTestSuite struct {
	suite.Suite
}

func (suite *RoomValueObjectTestSuite) SetupTest() {}

func (suite *NameValueObjectTestSuite) TestRoomWithInvalidLength() {
	room, err := valueobj.WithInvalidRoomLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewRoom",
		What:  "room must be between " + "1" + " to " + "64" + " characters",
		Why: errors.Meta{
			"Room": room,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *NameValueObjectTestSuite) TestRoomWithInvalidAlphanumeric() {
	room, err := valueobj.WithInvalidRoomAlphanumeric()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewRoom",
		What:  "room must be between " + "1" + " to " + "64" + " characters",
		Why: errors.Meta{
			"Room": room,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitRoomValueObjectSuite(t *testing.T) {
	suite.Run(t, new(RoomValueObjectTestSuite))
}

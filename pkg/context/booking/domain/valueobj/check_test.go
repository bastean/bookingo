package valueobj_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type CheckValueObjectTestSuite struct {
	suite.Suite
}

func (suite *CheckValueObjectTestSuite) SetupTest() {}

func (suite *CheckValueObjectTestSuite) TestCheck() {
	check, err := valueobj.InvalidCheck()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCheck",
		What:  "invalid check format, required e.g. DD-MM-YYYY",
		Why: errors.Meta{
			"Check": check,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCheckValueObjectSuite(t *testing.T) {
	suite.Run(t, new(CheckValueObjectTestSuite))
}

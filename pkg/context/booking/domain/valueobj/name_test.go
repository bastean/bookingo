package valueobj_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type NameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *NameValueObjectTestSuite) SetupTest() {}

func (suite *NameValueObjectTestSuite) TestNameWithInvalidLength() {
	name, err := valueobj.WithInvalidNameLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewName",
		What:  "name must be between " + "1" + " to " + "64" + " characters",
		Why: errors.Meta{
			"Name": name,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *NameValueObjectTestSuite) TestNameWithInvalidAlphanumeric() {
	name, err := valueobj.WithInvalidNameAlphanumeric()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewName",
		What:  "name must be between " + "1" + " to " + "64" + " characters",
		Why: errors.Meta{
			"Name": name,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitNameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(NameValueObjectTestSuite))
}

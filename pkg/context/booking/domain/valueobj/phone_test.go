package valueobj_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type PhoneValueObjectTestSuite struct {
	suite.Suite
}

func (suite *PhoneValueObjectTestSuite) SetupTest() {}

func (suite *PhoneValueObjectTestSuite) TestPhone() {
	phone, err := valueobj.InvalidPhone()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewPhone",
		What:  "invalid phone format, required e.g. +1123456789",
		Why: errors.Meta{
			"Phone": phone,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitPhoneValueObjectSuite(t *testing.T) {
	suite.Run(t, new(PhoneValueObjectTestSuite))
}

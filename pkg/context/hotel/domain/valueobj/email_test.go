package valueobj_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type EmailValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EmailValueObjectTestSuite) SetupTest() {}

func (suite *EmailValueObjectTestSuite) TestEmail() {
	email, err := valueobj.InvalidEmail()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEmail",
		What:  "invalid email format",
		Why: errors.Meta{
			"Email": email,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}

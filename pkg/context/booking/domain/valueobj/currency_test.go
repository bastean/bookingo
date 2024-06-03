package valueobj_test

import (
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/stretchr/testify/suite"
)

type CurrencyValueObjectTestSuite struct {
	suite.Suite
}

func (suite *CurrencyValueObjectTestSuite) SetupTest() {}

func (suite *CurrencyValueObjectTestSuite) TestCurrency() {
	currency, err := valueobj.InvalidCurrency()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCurrency",
		What:  "invalid currency format, required e.g. USD",
		Why: errors.Meta{
			"Currency": currency,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCurrencyValueObjectSuite(t *testing.T) {
	suite.Run(t, new(CurrencyValueObjectTestSuite))
}

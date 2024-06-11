package mail_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/communication/mail"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/transports"
	"github.com/stretchr/testify/suite"
)

type HotelMailAccountConfirmationTransportTestSuite struct {
	suite.Suite
	sut  models.Transport
	smtp *transports.SMTP
}

func (suite *HotelMailAccountConfirmationTransportTestSuite) SetupTest() {
	suite.smtp = transports.NewSMTP(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("URL"),
	)

	suite.sut = &mail.AccountConfirmation{
		SMTP: suite.smtp,
	}
}

func (suite *HotelMailAccountConfirmationTransportTestSuite) TestSubmit() {
	message := event.RandomCreatedSucceeded()

	attributes := new(event.CreatedSucceededAttributes)

	json.Unmarshal(message.Attributes, attributes)

	suite.NoError(suite.sut.Submit(attributes))
}

func TestIntegrationHotelMailAccountConfirmationTransportSuite(t *testing.T) {
	suite.Run(t, new(HotelMailAccountConfirmationTransportTestSuite))
}

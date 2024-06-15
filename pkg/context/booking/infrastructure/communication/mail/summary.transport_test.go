package mail_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/event"
	"github.com/bastean/bookingo/pkg/context/booking/infrastructure/communication/mail"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/transports"
	"github.com/stretchr/testify/suite"
)

type BookingMailSummaryTransportTestSuite struct {
	suite.Suite
	sut  models.Transport
	smtp *transports.SMTP
}

func (suite *BookingMailSummaryTransportTestSuite) SetupTest() {
	suite.smtp = transports.NewSMTP(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("URL"),
	)

	suite.sut = &mail.Summary{
		SMTP: suite.smtp,
	}
}

func (suite *BookingMailSummaryTransportTestSuite) TestSubmit() {
	message := event.RandomCreatedSucceeded()

	booking := new(event.CreatedSucceededAttributes)

	json.Unmarshal(message.Attributes, booking)

	suite.NoError(suite.sut.Submit(booking))
}

func TestIntegrationBookingMailSummaryTransportSuite(t *testing.T) {
	suite.Run(t, new(BookingMailSummaryTransportTestSuite))
}
